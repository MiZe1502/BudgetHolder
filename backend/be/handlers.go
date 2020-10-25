package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
				env.logger.Info("logging middleware")

				// logging request
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

func createMessageHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
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
	http.Handle("/message", middlewareChain.Then(http.HandlerFunc(createMessageHandler(env))))

	http.ListenAndServe(":8080", nil)

}
