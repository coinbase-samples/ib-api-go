/**
 * Copyright 2022-present Coinbase Global, Inc.
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
	"sync"

	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Redis      *redis.ClusterClient
	LogEntry   log.Entry
	Wait       *sync.WaitGroup
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
