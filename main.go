package main

import (
	"log"
	"net/http"

	"github.com/FilipeJohansson/gosocket"
)

func main() {
	handler := gosocket.NewHandler()
	handler.
		OnConnect(func(c *gosocket.Client) error {
			log.Println("Client connected")
			return nil
		}).
		OnMessage(func(c *gosocket.Client, m *gosocket.Message) error {
			log.Println("Message received:", string(m.RawData))
			c.SendMessage(m)
			return nil
		}).OnDisconnect(func(c *gosocket.Client) error {
		log.Println("Client disconnected")
		return nil
	})

	http.Handle("/ws", handler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
