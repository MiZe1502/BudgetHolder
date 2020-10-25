package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Hub stores current connections to clients
type Hub struct {
	connections map[*Connection]bool

	register   chan *Connection
	unregister chan *Connection
	send       chan []byte
}

// Connection represents single connection
type Connection struct {
	connectionData *ConnectionData
	conn           *websocket.Conn

	send chan []byte
}

// ConnectionData stores additional data about clients
// associated with this connection
type ConnectionData struct {
	ip      string
	session string
}

func createHub() *Hub {
	return &Hub{
		register:    make(chan *Connection),
		unregister:  make(chan *Connection),
		connections: make(map[*Connection]bool),
		send:        make(chan []byte),
	}
}

func (h *Hub) runHub() {
	for {
		select {
		case connection := <-h.register:
			h.connections[connection] = true
			fmt.Println(len(h.connections))
		case connection := <-h.unregister:
			if _, ok := h.connections[connection]; ok {
				delete(h.connections, connection)
				connection.conn.Close()
			}
		case message := <-h.send:
			fmt.Printf("message in hub: %s\n", string(message))
			for connection := range h.connections {
				err := connection.conn.WriteMessage(1, message)
				if err != nil {
					fmt.Println(err)
				}
				// select {
				// case connection.send <- message:
				// 	err := connection.conn.WriteMessage(1, message)
				// 	if err != nil {
				// 		fmt.Println(err)
				// 	}
				// 	// default:
				// 	// 	close(connection.send)
				// 	// 	delete(h.connections, connection)
				// }
			}
		}
	}
}
