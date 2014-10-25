package main

import (
	"log"
)

type Hub struct {
	// Registered connections.
	connections map[*Connection]bool
	// Inbound messages from the connections.
	broadcast chan Message
	// Register requests from the connections.
	register chan *Connection
	// Unregister requests from connections.
	unregister chan *Connection
}

var h = Hub{
	broadcast:   make(chan Message),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

func (h *Hub) run() {
	for {
		select {
		case c := <-h.register:
			log.Println("Register")
			h.connections[c] = true

			for _, m := range lastMessages {
				c.send <- m.ToJson()
			}
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				log.Println("Unregister")
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			// if current message >= maxLastMessages = delete first item from array
			if len(lastMessages) >= maxLastMessages {
				lastMessages = append(lastMessages[:0], lastMessages[1:]...)
			}
			// Add new message to lastMessages
			lastMessages = append(lastMessages, m)

			log.Println("Broadcast")
			for c := range h.connections {
				select {
				case c.send <- m.ToJson():
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}
