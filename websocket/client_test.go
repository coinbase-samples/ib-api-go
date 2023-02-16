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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/gorilla/websocket"
)

type testClientHandler struct{}

func (h testClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrade(w, r)
	if err != nil {
		return
	}
	cfg := config.AppConfig{
		BaseConfig: config.BaseConfig{
			Env: "local",
		},
	}
	pool := NewPool(cfg)

	go pool.Start()

	client := &Client{
		Conn:  conn,
		Pool:  pool,
		Alias: "Bob",
	}

	pool.Wait.Add(1)
	pool.Register <- client
	pool.Wait.Wait()

	pool.Wait.Add(1)
	client.Read()
	pool.Wait.Wait()

	pool.Wait.Add(1)
	pool.Unregister <- client
	pool.Wait.Wait()
}

func TestNewClient(t *testing.T) {
	h := &testClientHandler{}
	s := httptest.NewServer(h)
	wsURL := httpToWS(t, s.URL)

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	defer ws.Close()

}
