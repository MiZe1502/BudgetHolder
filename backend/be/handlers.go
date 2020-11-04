package main

import (
	"context"
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
	Trace        MiddlewareKey = "Trace"
)

var endpointsWithoutAuth = []string{"/api/v1/user/new", "/api/v1/user/auth"}

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
	case Trace:
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				traceUUID := utils.GetNewUUID()

				ctx := context.WithValue(r.Context(), "traceUUID", traceUUID)
				r = r.WithContext(ctx)

				env.logger.SetTraceUUID(traceUUID)
				env.logger.Info("init tracelog uuid: " + traceUUID)

				next.ServeHTTP(w, r)
			})
		}
	case CheckSession:
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				env.logger.Info("check user middleware")

				//check endpoints without auth
				requestPath := r.URL.Path

				env.logger.Info("requestPath: " + requestPath)

				for _, value := range endpointsWithoutAuth {
					if value == requestPath {
						next.ServeHTTP(w, r)
						return
					}
				}

				//get and parse token
				msg := make(map[string]interface{})
				tokenHeader := r.Header.Get("Authorization")

				env.logger.Info("check user token: " + tokenHeader)

				if tokenHeader == "" {
					msg = utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusBadRequest)
					utils.RespondError(w, msg, env.logger)
					return
				}

				token, err := env.token.ParseToken(tokenHeader)
				if err != nil {
					msg = utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
					utils.RespondError(w, msg, env.logger)
					return
				}

				env.logger.Info("token: " + tokenHeader + " is valid for session: " + token.SessionID.String())

				//get userID by sessionUUID from token
				var repo repos.UserRepository

				repo.SetDb(env.db)
				repo.SetLogger(env.logger)

				userID, err := repo.GetUserIDBySessionUUID(token.SessionID)
				if err != nil {
					msg = utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
					utils.RespondError(w, msg, env.logger)
					return
				}

				env.logger.Info("user with id: " + fmt.Sprint(userID) + " found for session: " + token.SessionID.String())

				//get user group id
				groupID, err := repo.GetUserGroupIDByUserID(userID)
				if err != nil {
					msg = utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
					utils.RespondError(w, msg, env.logger)
					return
				}

				//pass user data as context value
				userCtx := &repos.UserContext{
					SessionUUID: token.SessionID,
					UserID:      userID,
					UserGroupID: groupID,
					IP:          r.RemoteAddr,
				}

				env.logger.Info("group with id: " + fmt.Sprint(groupID) + " found for user with id: " + fmt.Sprint(userID))

				ctx := context.WithValue(r.Context(), "userCtx", userCtx)
				r = r.WithContext(ctx)

				next.ServeHTTP(w, r)
			})
		}
	}

	return nil
}

func createWsHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.logger.Info("createWsHandler")

		msg := make(map[string]interface{})

		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			msg = utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		userCtx := r.Context().Value("userCtx").(*repos.UserContext)
		traceUUID := r.Context().Value("traceUUID").(string)

		env.logger.Info("got context data for connection in hub")

		connectionData := &ConnectionData{userCtx: userCtx, traceUUID: traceUUID}
		connection := &Connection{conn: conn, connectionData: connectionData}

		env.hub.register <- connection
		env.logger.Info("registered connection in hub for user: " + fmt.Sprint(userCtx.UserID) + " with session: " + userCtx.SessionUUID.String())

		msgType, message, err := conn.ReadMessage()
		if err != nil {
			msg = utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		if err = conn.WriteMessage(msgType, message); err != nil {
			msg = utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}
	}
}

func createNewUserHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.logger.Info("createNewUserHandler")

		env.logger.Info("check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("getting user data from request")

		userData := &repos.FullUser{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("adding user: login: " + userData.Login)

		var repo repos.UserRepository

		repo.SetDb(env.db)
		repo.SetLogger(env.logger)
		repo.SetTokenGenerator(env.token)

		newUserID, err := repo.CreateNewUser(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("added user: login: " + userData.Login + " | id: " + fmt.Sprint(newUserID))

		userDataJSON, err := json.Marshal(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		utils.Respond(w, utils.MessageData(utils.Message(true, ""), userDataJSON), env.logger)
	}
}

func createAuthHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.logger.Info("createAuthHandler")

		env.logger.Info("check request method: " + r.Method)

		if r.Method != "POST" {
			msg := utils.MessageError(utils.Message(false, "Incorrect request method: "+r.Method), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("getting user data from request")

		userData := &repos.User{}
		err := json.NewDecoder(r.Body).Decode(userData)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, "Invalid request body"), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("processing user auth: login: " + userData.Login + " | pass: " + userData.Password)

		var repo repos.UserRepository

		repo.SetDb(env.db)
		repo.SetLogger(env.logger)
		repo.SetTokenGenerator(env.token)

		token, err := repo.ProcessUserAuth(userData.Login, userData.Password, r.RemoteAddr)
		if err != nil {
			msg := utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.logger)
			return
		}

		env.logger.Info("got token: " + token + " for user: " + userData.Login)

		utils.Respond(w, utils.Message(true, token), env.logger)
	}
}

func createTestMessageHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		user := r.Context().Value("userCtx").(*repos.UserContext)

		// Print the message to the console
		fmt.Printf("message: %s\n", fmt.Sprint(user.UserID))

		env.hub.send <- body
	}
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.html")
}

func initHandlers(env *Env) {
	middlewareChain := alice.New(createMiddleware(env, Trace),
		createMiddleware(env, Logger),
		createMiddleware(env, CheckSession))

	http.Handle("/", middlewareChain.Then(http.HandlerFunc(serveStatic)))
	http.Handle("/ws", middlewareChain.Then(http.HandlerFunc(createWsHandler(env))))

	http.Handle("/api/v1/user/auth", middlewareChain.Then(http.HandlerFunc(createAuthHandler(env))))
	http.Handle("/api/v1/user/new", middlewareChain.Then(http.HandlerFunc(createNewUserHandler(env))))

	http.Handle("/message", middlewareChain.Then(http.HandlerFunc(createTestMessageHandler(env))))

	http.ListenAndServe(":8080", nil)

}
