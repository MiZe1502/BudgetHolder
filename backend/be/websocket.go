package main

import (
	"fmt"
	"net/http"

	repos "./repositories"
	utils "./utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
