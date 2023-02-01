package websocket

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

type testHandler struct{}

func TestUpgrade(t *testing.T) {
	h := &testHandler{}
	s := httptest.NewServer(h)
	wsURL := httpToWS(t, s.URL)

	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	defer ws.Close()
}

func (h testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := Upgrade(w, r)
	if err != nil {
		return
	}
}

func httpToWS(t *testing.T, u string) string {
	t.Helper()

	wsURL, err := url.Parse(u)
	if err != nil {
		t.Fatal(err)
	}

	switch wsURL.Scheme {
	case "http":
		wsURL.Scheme = "ws"
	case "https":
		wsURL.Scheme = "wss"
	}

	return wsURL.String()
}
