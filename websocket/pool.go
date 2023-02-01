package websocket

import (
	"crypto/tls"
	"fmt"
	"sync"

	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/go-redis/redis"
)

func NewPool(conf config.AppConfig) *Pool {
	redisClient := makeClient(conf)

	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Redis:      redisClient,
		Wait:       &sync.WaitGroup{},
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Debugf("registered new client: %s", client.Alias)
			pool.Wait.Done()
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Debugf("Unregistered client: %s, new Size of Connection Pool: %d", client.ID, len(pool.Clients))
			pool.Wait.Done()
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
