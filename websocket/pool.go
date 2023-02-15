/**
 * Copyright 2022 - Present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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

	clusterOpts := &redis.ClusterOptions{
		Addrs: addrs,
	}

	if !conf.IsLocalEnv() {
		clusterOpts.TLSConfig = &tls.Config{}
	}

	return redis.NewClusterClient(clusterOpts)
}
