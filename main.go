package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Configures the WebSocket connection upgrade
// CheckOrigin a function that allows connections from any origin
// For security reason might use here
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handles WebSocket connections.
// upgrader.Upgrade() : Method upgrades the incoming HTTP request to a WebSocket
// connection. It takes the http.ResponseWriter, the *http.Request,
// and an optional response header as parameters. If the upgrade is successful,
// it returns a *websocket.Conn representing the WebSocket connection
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)

	// If there is an error during the upgrade process, it is captured in the err variable.
	if err != nil {
		log.Fatal(err)
		return
	}
	//Ensures that the WebSocket connection is closed when the function exits.
	defer conn.Close()

	// Infinite loop to handle WebSocket messages
	for {
		//conn.ReadMessage(): Reads a message from the WebSocket connection.
		//It returns the message type, the message payload (p), and any error that may occur during the process.
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			//if an error occurs the function logs the error and exits the loop.
			return
		}

		// Print the received message to the console vscode
		fmt.Printf("Received message: %s\n", p)

		// display/echo the messege
		// conn.WriteMessage : the server echoes the same message back to the client after receiving the message.
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Route requests to /ws path to the handleConnections function above
	http.HandleFunc("/ws", handleConnections)

	// Run the server on port 8080
	err := http.ListenAndServe(":8080", nil)
	//if error
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
