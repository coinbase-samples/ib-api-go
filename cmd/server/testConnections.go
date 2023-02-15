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
	"time"

	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/log"
	asset "github.com/coinbase-samples/ib-api-go/pkg/pbs/asset/v1"
	profile "github.com/coinbase-samples/ib-api-go/pkg/pbs/profile/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// standalone dial function to verify connectivity in different environments (local, aws, in docker)
func testProfileDial(app config.AppConfig) {
	dialProfileConn := getProfileConnAddress(app)
	clientCreds := getGrpcCredentials(app)
	grpc.EnableTracing = true

	conn, err := grpc.Dial(dialProfileConn, grpc.WithTransportCredentials(clientCreds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := profile.NewProfileServiceClient(conn)

	var testUserId = "37d10e18-34a2-4bd2-b7bc-b8e6dd6358f1"
	if app.IsLocalEnv() {
		testUserId = "c7e34d37-f678-4096-94f7-3cad7d3258b9"
	}

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	md := metadata.New(map[string]string{"x-route-id": app.UserRouteId})
	ctx = metadata.NewOutgoingContext(ctx, md)
	log.Debugf("dialing profile with - %s - %s - %v", dialProfileConn, testUserId, ctx)
	r, err := c.ReadProfile(ctx, &profile.ReadProfileRequest{Id: testUserId})
	grpc.EnableTracing = false

	if err != nil {
		log.Warnf("could not greet profile: %v", err)
		return
	}
	log.Warnf("Greeting Profile: %s", r.UserName)
}

func testOrderDial(app config.AppConfig) {
	dialOrderConn := getOrderConnAddress(app)
	clientCreds := getGrpcCredentials(app)
	grpc.EnableTracing = true

	conn, err := grpc.Dial(dialOrderConn, grpc.WithTransportCredentials(clientCreds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := asset.NewAssetServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	md := metadata.New(map[string]string{"x-route-id": app.OrderRouteId})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	log.Warnf("sending order test %s - %v - %v", dialOrderConn, ctx, md)
	r, err := c.ListAssets(ctx, &asset.ListAssetsRequest{})
	grpc.EnableTracing = false

	if err != nil {
		log.Warnf("could not greet order: %v", err)
		return
	}
	log.Warnf("Greeting Order: %s", r.Data)
}
