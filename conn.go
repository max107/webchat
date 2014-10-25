package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type Connection struct {
	// The websocket connection.
	ws *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *Connection) wsReader() {
	for {
		msg := Message{}
		err := c.ws.ReadJSON(&msg)
		if err != nil {
			break
		}
		h.broadcast <- msg
	}
	c.ws.Close()
}

func (c *Connection) wsWriter() {
	for msg := range c.send {
		log.Printf("%s", msg)
		err := c.ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
