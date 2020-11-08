package main

import (
	"context"
	"fmt"
	"net/http"

	repos "./repositories"
	utils "./utils"
)

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