package websocket

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

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
