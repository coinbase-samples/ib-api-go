package websocket

import (
	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	Redis      *redis.ClusterClient
	LogEntry   log.Entry
}

type Client struct {
	ID            string
	Conn          *websocket.Conn
	Pool          *Pool
	Alias         string
	Subscriptions []*redis.PubSub
}

type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
}
