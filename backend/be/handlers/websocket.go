package handlers

import (
	"fmt"
	"net/http"

	repos "../repositories"
	utils "../utils"
	env "../env"
	wshub "../hub"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func createWsHandler(env *env.Env, hub *wshub.Hub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Logger.Info("createWsHandler")

		msg := make(map[string]interface{})

		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			msg = utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		userCtx := r.Context().Value("userCtx").(*repos.UserContext)
		traceUUID := r.Context().Value("traceUUID").(string)

		env.Logger.Info("got context data for connection in hub")

		connectionData := &wshub.ConnectionData{UserCtx: userCtx, TraceUUID: traceUUID}
		connection := &wshub.Connection{Conn: conn, ConnectionData: connectionData}

		hub.Register <- connection
		env.Logger.Info("registered connection in hub for user: " + fmt.Sprint(userCtx.UserID) + " with session: " + userCtx.SessionUUID.String())

		msgType, message, err := conn.ReadMessage()
		if err != nil {
			msg = utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}

		if err = conn.WriteMessage(msgType, message); err != nil {
			msg = utils.MessageError(utils.Message(false, err.Error()), http.StatusInternalServerError)
			utils.RespondError(w, msg, env.Logger)
			return
		}
	}
}
