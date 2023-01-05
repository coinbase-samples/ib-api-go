package websocket

import (
	"encoding/json"

	"github.com/coinbase-samples/ib-api-go/log"
)

func (c *Client) Read() {
	defer func() {
		for _, sub := range c.Subscriptions {
			sub.Close()
		}
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Errorf("read message error: %v", err)
			return
		}

		var message Message
		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Errorf("unmarshal error: - %v", err)
		}
		log.Debugf("%+v", message)
	}
}
