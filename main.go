package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error doing upgrade:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")
	for {
		messageType, message, err := conn.ReadMessage()
		if err!= nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Println("Message received:", string(message))

		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Error to write message:", err)
			break
		}
	}
	log.Println("Client disconnected")
}

func main() {
	http.HandleFunc("/ws", handler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
