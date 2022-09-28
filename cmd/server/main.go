package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/dba"
	log "github.com/sirupsen/logrus"
)

var (
	//setup logrus for interceptor
	logrusLogger = log.New()
	wait         time.Duration
)

func main() {
	var app config.AppConfig

	config.Setup(&app)
	fmt.Println("starting app with config", app)

	config.LogInit(app, logrusLogger)

	//setup database
	repo := dba.NewRepo(&app)
	dba.NewDBA(repo)

	gwServer, err := setupHttp(app)
	if err != nil {
		logrusLogger.Warnln("issues setting up http server", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if gwServer != nil {
		gwServer.Shutdown(ctx)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logrusLogger.Debugln("stopping")
	os.Exit(0)
}
