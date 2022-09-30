package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID    string
	Conn  *websocket.Conn
	Pool  *Pool
	Alias string
}

type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var message Message
		err = json.Unmarshal(p, &message)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%+v", message)
	}
}
