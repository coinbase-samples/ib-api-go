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
	"encoding/json"

	"github.com/coinbase-samples/ib-api-go/log"
)

func (c *Client) Read() {
	defer func() {
		for _, sub := range c.Subscriptions {
			sub.Close()
		}
		c.Pool.Wait.Add(1)
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
