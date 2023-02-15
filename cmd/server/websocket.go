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

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/coinbase-samples/ib-api-go/dba"
	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/coinbase-samples/ib-api-go/websocket"
	"github.com/go-redis/redis"
)

func serveWs(ctx context.Context, pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.DebugCtx(ctx, "WebSocket Endpoint Hit", r.URL.Query())

	query := r.URL.Query()
	alias := query.Get("alias")
	if alias == "" {
		log.DebugCtx(ctx, "Missing required connection params", query, alias)
		return
	}

	log.DebugfCtx(ctx, "adding new ws connection - %s", alias)
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.WarnfCtx(ctx, "%+v\n%v", w, err)
		return
	}

	orderChannelName := fmt.Sprintf("%s-orders", alias)
	log.DebugfCtx(ctx, "starting subscription - %v", orderChannelName)
	orderSub := pool.Redis.Subscribe(orderChannelName)
	defer orderSub.Close()
	ch := orderSub.Channel()

	go func() {
		for msg := range ch {
			mess := websocket.Message{}
			if err := json.Unmarshal([]byte(msg.Payload), &mess); err != nil {
				log.WarnfCtx(ctx, "error marshalling order status message - %v", err)
				return
			}

			log.DebugfCtx(ctx, "order sub message - %v - %v", alias, mess)
			conn.WriteJSON(mess)
		}
	}()

	client := &websocket.Client{
		Conn:          conn,
		Pool:          pool,
		Alias:         alias,
		Subscriptions: []*redis.PubSub{orderSub},
	}

	//publish initial open/pending orders
	checkOrdersFromDynamo(ctx, client)
	pool.Wait.Add(2)
	pool.Register <- client
	client.Read()
}

func assetPriceUpdater(pool websocket.Pool) {
	// every second send new asset prices to current clients
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

				log.Debugf("asset price loop for clients: %d", len(pool.Clients))
				if len(pool.Clients) > 0 {
					assets, err := dba.Repo.ListAssets(context.Background(), "")
					if err != nil {
						log.Warnf("error reading assets for price update - %v", err)
					} else {
						body, _ := json.Marshal(assets)
						message := websocket.Message{Type: "assets", Body: string(body)}
						for client := range pool.Clients {
							client.Conn.WriteJSON(message)
						}
					}

				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func checkOrdersFromDynamo(ctx context.Context, client *websocket.Client) {
	orders, err := dba.Repo.ListOrders(context.Background(), client.Alias)
	if err != nil {
		log.WarnfCtx(ctx, "error reading orders for price update - %v", err)
	} else {
		if len(orders) < 1 {
			log.DebugfCtx(ctx, "skipping order update -%v", orders)
		} else {
			body, err := json.Marshal(orders)
			if err != nil {
				log.WarnfCtx(ctx, "issue marshalling existing orders - %v", err)
			}
			message := websocket.Message{Type: "orders", Body: string(body)}
			log.DebugfCtx(ctx, "writing initial order status - %v", message)
			client.Conn.WriteJSON(message)
		}
	}

}
