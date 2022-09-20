package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/coinbase-samples/ib-api-go/config"
	v1 "github.com/coinbase-samples/ib-api-go/pkg/pbs/v1"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func orderConn(app config.AppConfig) (*grpc.ClientConn, error) {
	var dialOrderConn string
	if app.Env == "local" {
		dialOrderConn = fmt.Sprintf("%s:%s", app.NetworkName, app.OrderGrpcPort)
	} else {
		dialOrderConn = fmt.Sprintf("%s/orders", app.NetworkName)
	}
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		dialOrderConn,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return conn, err
}

func profileConn(app config.AppConfig) (*grpc.ClientConn, error) {
	var dialProfileConn string
	if app.Env == "local" {
		dialProfileConn = fmt.Sprintf("%s:%s", app.NetworkName, app.GrpcPort)
	} else {
		dialProfileConn = fmt.Sprintf("%s/orders", app.NetworkName)
	}
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		dialProfileConn,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return conn, err
}

func setupHttp(app config.AppConfig) (*http.Server, error) {
	logrusLogger.Debugln("dialing order manager")
	oConn, err := orderConn(app)
	if err != nil {
		logrusLogger.Fatalln("Failed to dial server:", err)
	}
	logrusLogger.Debugln("Connected to order manager")

	logrusLogger.Debugln("dialing profile")
	pConn, err := profileConn(app)
	if err != nil {
		logrusLogger.Fatalln("Failed to dial server:", err)
	}
	logrusLogger.Debugln("Connected to profile")

	gwmux := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
		md := make(map[string]string)
		if method, ok := runtime.RPCMethod(ctx); ok {
			md["method"] = method // /grpc.gateway.examples.internal.proto.examplepb.LoginService/Login
		}
		if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
			md["pattern"] = pattern // /v1/example/login
		}
		return metadata.New(md)
	}))

	// Register Service Handlers
	err = v1.RegisterOrderServiceHandler(context.Background(), gwmux, oConn)
	if err != nil {
		logrusLogger.Fatalln("Failed to register order:", err)
	}
	err = v1.RegisterProfileServiceHandler(context.Background(), gwmux, pConn)
	if err != nil {
		logrusLogger.Fatalln("Failed to register profile:", err)
	}
	err = v1.RegisterBalanceServiceHandler(context.Background(), gwmux, oConn)
	if err != nil {
		logrusLogger.Fatalln("Failed to register balance:", err)
	}
	err = v1.RegisterAssetServiceHandler(context.Background(), gwmux, oConn)
	if err != nil {
		logrusLogger.Fatalln("Failed to register asset:", err)
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{
		"https://api.neoworks.dev",
		fmt.Sprintf("https://localhost:%s", app.Port),
		fmt.Sprintf("http://localhost:%s", app.Port),
		"http://localhost:4200",
		"https://localhost:4200",
	})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	gwServer := &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(gwmux),
		Addr:         fmt.Sprintf(":%s", app.Port),
		WriteTimeout: 40 * time.Second,
		ReadTimeout:  40 * time.Second,
	}

	logrusLogger.Println("Serving gRPC-Gateway on:", app.NetworkName, app.Port)

	go func() {
		if app.Env == "production" {
			if err := gwServer.ListenAndServeTLS("server.crt", "server.key"); err != nil {
				logrusLogger.Fatalln("ListenAndServeTLS: ", err)
			}
		} else {
			if err := gwServer.ListenAndServe(); err != nil {
				logrusLogger.Fatalln("ListenAndServe: ", err)
			}
		}
	}()

	return gwServer, nil
}
