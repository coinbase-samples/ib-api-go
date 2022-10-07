package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/coinbase-samples/ib-api-go/dba"
	"github.com/coinbase-samples/ib-api-go/websocket"
	"github.com/go-redis/redis"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	logrusLogger.Debugln("WebSocket Endpoint Hit", r.URL.Query())

	query := r.URL.Query()
	alias := query.Get("alias")
	if alias == "" {
		logrusLogger.Debugln("Missing required connection params", query, alias)
		return
	}

	logrusLogger.Debugf("adding new ws connection - %s", alias)
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	orderChannelName := fmt.Sprintf("%s-orders", alias)
	logrusLogger.Debugf("starting subscription - %v", orderChannelName)
	orderSub := pool.Redis.Subscribe(orderChannelName)
	defer orderSub.Close()
	ch := orderSub.Channel()

	go func() {
		for msg := range ch {
			mess := websocket.Message{}
			if err := json.Unmarshal([]byte(msg.Payload), &mess); err != nil {
				logrusLogger.Warnf("error marshalling order status message - %v", err)
				return
			}

			logrusLogger.Debugf("order sub message - %v - %v", alias, mess)
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
	checkOrdersFromDynamo(client)

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

				logrusLogger.Debugf("asset price loop for clients: %d", len(pool.Clients))
				if len(pool.Clients) > 0 {
					assets, err := dba.Repo.ListAssets(context.Background(), "")
					if err != nil {
						logrusLogger.Warnf("error reading assets for price update - %v", err)
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

func checkOrdersFromDynamo(client *websocket.Client) {

	orders, err := dba.Repo.ListOrders(context.Background(), client.Alias)
	if err != nil {
		logrusLogger.Warnf("error reading orders for price update - %v", err)
	} else {
		if len(orders) < 1 {
			logrusLogger.Debugf("skipping order update -%v", orders)
		} else {
			body, err := json.Marshal(orders)
			if err != nil {
				logrusLogger.Warnf("issue marshalling existing orders - %v", err)
			}
			message := websocket.Message{Type: "orders", Body: string(body)}
			logrusLogger.Debugf("writing initial order status - %v", message)
			client.Conn.WriteJSON(message)
		}
	}

}
