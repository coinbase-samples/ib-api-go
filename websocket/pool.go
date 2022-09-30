package websocket

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true

		case client := <-pool.Unregister:

			delete(pool.Clients, client)
			fmt.Println("Unregistered client, new Size of Connection Pool: ", len(pool.Clients))

		// I don't see a use case for broadcast between clients. but leaving in case
		case message := <-pool.Broadcast:
			///Send out updated
			for client, _ := range pool.Clients {
				fmt.Println("Should send to client?", client.Alias, message)
			}
		}
	}

}
