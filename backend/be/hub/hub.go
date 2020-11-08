package hub

import (
	"fmt"

	env "../env"
	repos "../repositories"

	"github.com/gorilla/websocket"
)

// Hub stores current connections to clients
type Hub struct {
	connections map[*Connection]bool

	Register   chan *Connection
	unregister chan *Connection
	Send       chan []byte
}

// Connection represents single connection
type Connection struct {
	ConnectionData *ConnectionData
	Conn           *websocket.Conn

	Send chan []byte
}

// ConnectionData stores additional data about clients
// associated with this connection
type ConnectionData struct {
	UserCtx   *repos.UserContext
	TraceUUID string
}

// CreateHub creates messages hub instance
func CreateHub() *Hub {
	return &Hub{
		Register:    make(chan *Connection),
		unregister:  make(chan *Connection),
		connections: make(map[*Connection]bool),
		Send:        make(chan []byte),
	}
}

// RunHub execute main hub logic in infinite loop
func (h *Hub) RunHub(env *env.Env) {
	for {
		select {
		case connection := <-h.Register:
			h.connections[connection] = true
			env.Logger.Info("connections registered: " + fmt.Sprint(len(h.connections)))
		case connection := <-h.unregister:
			if _, ok := h.connections[connection]; ok {
				delete(h.connections, connection)
				connection.Conn.Close()
			}
		case message := <-h.Send:
			env.Logger.Info("got message in hub: " + string(message))
			for connection := range h.connections {
				err := connection.Conn.WriteMessage(1, message)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

// CloseAllConnections destroys all connections stored in hub
func (h *Hub) CloseAllConnections(env *env.Env) {
	for connection := range h.connections {
		h.unregister <- connection
	}
}
