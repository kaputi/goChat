package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket  *websocket.Conn
	receive chan []byte
	room    *room
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
}

func (c *client) write() {
  defer c.socket.Close()

  for msg := range c.receive {
    if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
      break
    }
  }
}
