package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/coinbase-samples/ib-api-go/auth"
	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/dba"
	"github.com/coinbase-samples/ib-api-go/log"
)

var (
	wait time.Duration
)

func main() {
	var app config.AppConfig

	config.Setup(&app)
	fmt.Println("starting app with config", app)

	log.Init(app)

	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("error creating dynamo config: %v", err)

	}

	//setup auth client
	cip := auth.InitAuth(&app, cfg)
	aw := auth.Middleware{Cip: cip}

	//setup database
	repo := dba.NewRepo(&app, cfg)
	dba.NewDBA(repo)

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	gwServer, err := setupHttp(ctx, app, aw)
	if err != nil {
		log.Warn("issues setting up http server", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	if gwServer != nil {
		gwServer.Shutdown(ctx)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Debug("stopping")
	os.Exit(0)
}
