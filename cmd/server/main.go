/**
 * Copyright 2022 - Present Coinbase Global, Inc.
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

package main

import (
	"context"
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

	log.Init(app)
	log.Debugf("starting app with config - %v", app)

	cfg, err := awsConfig.LoadDefaultConfig(context.Background())
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
