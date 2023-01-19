package websocket

import (
	"testing"

	"github.com/coinbase-samples/ib-api-go/config"
)

func TestNewPool(t *testing.T) {
	cfg := config.AppConfig{
		BaseConfig: config.BaseConfig{
			Env: "local",
		},
	}
	pool := NewPool(cfg)

	if pool.Redis == nil {
		t.Fatal("expected redis client to be set")
	}

	if pool.Redis.Options().TLSConfig != nil {
		t.Fatal("expected TLS to not be set on local env")
	}
}

func TestNewPoolNotLocal(t *testing.T) {
	cfg := config.AppConfig{
		BaseConfig: config.BaseConfig{
			Env: "stage",
		},
	}
	pool := NewPool(cfg)

	if pool.Redis == nil {
		t.Fatal("expected redis client to be set")
	}

	if pool.Redis.Options().TLSConfig == nil {
		t.Fatal("expected TLS to be set on other envs")
	}
}

func TestPoolRegister(t *testing.T) {
	cfg := config.AppConfig{
		BaseConfig: config.BaseConfig{
			Env: "stage",
		},
	}
	pool := NewPool(cfg)

	go pool.Start()

	client := &Client{
		Pool:  pool,
		Alias: "bob",
	}

	pool.Wait.Add(1)
	pool.Register <- client
	pool.Wait.Wait()
	if len(pool.Clients) != 1 {
		t.Fatal("expected client added to the pool")
	}

	pool.Wait.Add(1)
	client2 := &Client{
		Pool:  pool,
		Alias: "alice",
	}

	pool.Register <- client2
	pool.Wait.Wait()
	if len(pool.Clients) != 2 {
		t.Fatal("expected client2 added to the pool", len(pool.Clients))
	}
}

func TestPoolUnregister(t *testing.T) {
	cfg := config.AppConfig{
		BaseConfig: config.BaseConfig{
			Env: "stage",
		},
	}
	pool := NewPool(cfg)

	go pool.Start()

	client := &Client{
		Pool:  pool,
		Alias: "bob",
	}

	pool.Wait.Add(1)
	pool.Register <- client
	pool.Wait.Wait()

	if len(pool.Clients) != 1 {
		t.Fatal("expected client added to the pool")
	}

	pool.Wait.Add(1)
	pool.Unregister <- client
	pool.Wait.Wait()
	if pool.Clients[client] {
		t.Fatal("expected client removed from the pool")
	}
}
