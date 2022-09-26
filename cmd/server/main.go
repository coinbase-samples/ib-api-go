package main

import (
	"fmt"
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
	logLevel, _ := log.ParseLevel("debug") //app.LogLevel)
	logrusLogger.SetLevel(logLevel)
	//setup cognito client
	cip := InitAuth(&app)
	aw := authMiddleware{cip} //setup dynamodb connection

	//setup database
	repo := dba.NewRepo(&app)
	dba.NewDBA(repo)

	// Start gRPC Server
	gRPCListen(app, aw)
}
