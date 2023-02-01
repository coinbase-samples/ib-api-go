package main

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/log"
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

func orderConn(app config.AppConfig, dialCreds credentials.TransportCredentials) (*grpc.ClientConn, error) {
	dialOrderConn := getOrderConnAddress(app)

	md := metadata.New(map[string]string{"x-route-id": app.OrderRouteId})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	log.Debug("order dial", dialOrderConn, md)
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests

	conn, err := grpc.DialContext(
		ctx,
		dialOrderConn,
		grpc.WithTransportCredentials(dialCreds),
	)
	return conn, err
}

func profileConn(app config.AppConfig, dialCreds credentials.TransportCredentials) (*grpc.ClientConn, error) {
	dialProfileConn := getProfileConnAddress(app)

	md := metadata.New(map[string]string{"x-route-id": app.UserRouteId})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	log.Debug("profile dial", dialProfileConn)
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		ctx,
		dialProfileConn,
		grpc.WithTransportCredentials(dialCreds),
	)
	return conn, err
}
