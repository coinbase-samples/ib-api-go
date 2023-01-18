package websocket

import (
	"net/http"

	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("error upgrading ws connection: %v", err)
		return nil, err
	}

	return conn, nil
}
