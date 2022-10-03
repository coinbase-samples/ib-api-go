package main

import (
	"context"
	"log"
	"time"

	"github.com/coinbase-samples/ib-api-go/config"
	profile "github.com/coinbase-samples/ib-api-go/pkg/pbs/profile/v1"
	v1 "github.com/coinbase-samples/ib-api-go/pkg/pbs/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

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

	var testId = "c5af3271-7185-4a52-9d0c-1c4b418317d8"
	if app.IsLocalEnv() {
		testId = "c7e34d37-f678-4096-94f7-3cad7d3258b9"
	}

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	md := metadata.New(map[string]string{"x-route-id": app.UserRouteId})
	ctx = metadata.NewOutgoingContext(ctx, md)
	logrusLogger.Debugf("dialing profile with - %s - %v", dialProfileConn, ctx, testId)
	r, err := c.ReadProfile(ctx, &profile.ReadProfileRequest{Id: testId})
	grpc.EnableTracing = false

	if err != nil {
		logrusLogger.Warnf("could not greet profile: %v", err)
		return
	}
	logrusLogger.Warnf("Greeting Profile: %s", r.UserName)
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

	c := v1.NewAssetServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	md := metadata.New(map[string]string{"x-route-id": app.OrderRouteId})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	logrusLogger.Warnf("sending order test %s - %v - %v", dialOrderConn, ctx, md)
	r, err := c.ListAssets(ctx, &v1.ListAssetsRequest{})
	grpc.EnableTracing = false

	if err != nil {
		logrusLogger.Warnf("could not greet order: %v", err)
		return
	}
	logrusLogger.Warnf("Greeting Order: %s", r.Data)
}
