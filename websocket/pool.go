package websocket

import (
	"crypto/tls"
	"fmt"

	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/go-redis/redis"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	Redis      *redis.ClusterClient
	LogEntry   log.Entry
}

func NewPool(conf config.AppConfig) *Pool {
	redisClient := makeClient(conf)

	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		Redis:      redisClient,
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true

		case client := <-pool.Unregister:

			delete(pool.Clients, client)
			log.Debugf("Unregistered client, new Size of Connection Pool: %d", len(pool.Clients))

		// I don't see a use case for broadcast between clients. but leaving in case
		case message := <-pool.Broadcast:
			///Send out updated
			for client := range pool.Clients {
				log.Debugf("Should send to client? %s - %v", client.Alias, message)
			}
		}
	}

}

func makeClient(conf config.AppConfig) *redis.ClusterClient {
	addrs := []string{fmt.Sprintf("%s:%s", conf.RedisEndpoint, conf.RedisPort)}

	if conf.IsLocalEnv() {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: addrs,
		})
	} else {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:     addrs,
			TLSConfig: &tls.Config{},
		})
	}
}
