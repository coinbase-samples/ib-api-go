package main

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/coinbase-samples/ib-api-go/config"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func getGrpcCredentials(app config.AppConfig) credentials.TransportCredentials {
	if app.IsLocalEnv() {
		return insecure.NewCredentials()
	} else {
		return credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})
	}
}

func getOrderConnAddress(app config.AppConfig) string {
	if app.IsLocalEnv() {
		return fmt.Sprintf("%s:%s", app.OrderMgrHostname, app.OrderGrpcPort)
	}
	return fmt.Sprintf("%s:443", app.OrderMgrHostname)
}

func getProfileConnAddress(app config.AppConfig) string {
	if app.IsLocalEnv() {
		return fmt.Sprintf("%s:%s", app.UserMgrHostname, app.UserGrpcPort)
	}
	return fmt.Sprintf("%s:443", app.UserMgrHostname)
}

func orderConn(app config.AppConfig) (*grpc.ClientConn, error) {
	dialOrderConn := getOrderConnAddress(app)
	clientCreds := getGrpcCredentials(app)

	md := metadata.New(map[string]string{"x-route-id": app.OrderRouteId})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	logrusLogger.Debugln("order dial", dialOrderConn, md)
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests

	conn, err := grpc.DialContext(
		ctx,
		dialOrderConn,
		grpc.WithTransportCredentials(clientCreds),
	)
	return conn, err
}

func profileConn(app config.AppConfig) (*grpc.ClientConn, error) {
	dialProfileConn := getProfileConnAddress(app)
	dialCreds := getGrpcCredentials(app)
	logrusLogger.Debugln("connecting to profile localhost grpc", dialProfileConn)
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		dialProfileConn,
		grpc.WithTransportCredentials(dialCreds),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	)
	return conn, err
}
