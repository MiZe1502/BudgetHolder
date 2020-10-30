package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	repos "./repositories"
	utils "./utils"
	"github.com/gorilla/websocket"
	"github.com/justinas/alice"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// MiddlewareKey is a type for middlewares keys enum
type MiddlewareKey string

const (
	Logger       MiddlewareKey = "Logger"
	CheckSession MiddlewareKey = "CheckSession"
)

func createMiddleware(env *Env, middleWareType MiddlewareKey) func(next http.Handler) http.Handler {
	switch middleWareType {
	case Logger:
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				env.logger.Info("got request")
				env.logger.Info("user agent: " + r.UserAgent() + " | " + "ip: " + r.RemoteAddr + " | " + "method: " + r.Method)
				next.ServeHTTP(w, r)
			})
		}
	case CheckSession:
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				env.logger.Info("check user middleware")
				// check user token and session
				next.ServeHTTP(w, r)
			})
		}
	}

	return nil
}

func createWsHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.logger.Info("create ws handler")

		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		connectionData := &ConnectionData{ip: r.RemoteAddr, session: utils.GetNewUUID()}
		connection := &Connection{conn: conn, connectionData: connectionData}

		env.hub.register <- connection
		env.logger.Info("registered connection")

		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func createAuthHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userData := &repos.User{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), utils.ServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		var repo repos.UserRepository

		repo.SetDb(env.db)
		repo.SetLogger(env.logger)
		repo.SetTokenGenerator(env.token)

		token, err := repo.ProcessUserAuth(userData.Login, userData.Password, r.RemoteAddr)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), utils.ServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		//create web socket connection

		utils.Respond(w, utils.Message(true, token), env.logger)
	}
}

func createTestMessageHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		// Print the message to the console
		fmt.Printf("message: %s\n", string(body))

		env.hub.send <- body
	}
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.html")
}

func initHandlers(env *Env) {
	middlewareChain := alice.New(createMiddleware(env, Logger), createMiddleware(env, CheckSession))

	http.Handle("/", middlewareChain.Then(http.HandlerFunc(serveStatic)))
	http.Handle("/ws", middlewareChain.Then(http.HandlerFunc(createWsHandler(env))))

	http.Handle("/api/v1/auth", middlewareChain.Then(http.HandlerFunc(createAuthHandler(env))))

	http.Handle("/message", middlewareChain.Then(http.HandlerFunc(createTestMessageHandler(env))))

	http.ListenAndServe(":8080", nil)

}
