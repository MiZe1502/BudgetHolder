package main

import (
	"fmt"
	"net/http"

	utils "./utils"
	"github.com/gorilla/websocket"
	"github.com/justinas/alice"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("logging middleware")
		// logging request
		next.ServeHTTP(w, r)
	})
}

func checkSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("check user")
		// check user token and session
		next.ServeHTTP(w, r)
	})
}

func createWsHandler(env *Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		connectionData := &ConnectionData{ip: r.RemoteAddr, session: utils.GetNewUUID()}
		connection := &Connection{conn: conn, connectionData: connectionData}

		env.hub.register <- connection

		fmt.Println(len(env.hub.connections))

		msgType, msg, err := conn.ReadMessage()

		fmt.Println(msgType)


		if err != nil {
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		env.hub.send <- msg

		// for {
		// 	// Read message from browser
		// 	msgType, msg, err := conn.ReadMessage()
		// 	if err != nil {
		// 		return
		// 	}

		// 	// Print the message to the console
		// 	fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// 	// Write message back to browser
		// 	if err = conn.WriteMessage(msgType, msg); err != nil {
		// 		return
		// 	}
		// }
	}

}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.html")
}

func initHandlers(env *Env) {
	middlewareChain := alice.New(loggerMiddleware, checkSessionMiddleware)

	http.Handle("/", middlewareChain.Then(http.HandlerFunc(serveStatic)))
	http.Handle("/echo", middlewareChain.Then(http.HandlerFunc(createWsHandler(env))))

	http.ListenAndServe(":8080", nil)

}
