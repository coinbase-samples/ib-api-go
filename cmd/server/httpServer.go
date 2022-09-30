package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"

	"github.com/coinbase-samples/ib-api-go/config"
	profile "github.com/coinbase-samples/ib-api-go/pkg/pbs/profile/v1"
	v1 "github.com/coinbase-samples/ib-api-go/pkg/pbs/v1"
	"github.com/coinbase-samples/ib-api-go/websocket"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

		if strings.HasPrefix(r.URL.String(), "/v1/profile") {
			md["x-route-id"] = app.UserRouteId
			logrusLogger.Debugln("/v1/profile adding profile route id", app.UserRouteId)
		} else if strings.HasPrefix(r.URL.String(), "/v1/order") {
			md["x-route-id"] = app.OrderRouteId
			logrusLogger.Debugln("/v1/order adding order route id", app.OrderRouteId)
		} else if strings.HasPrefix(r.URL.String(), "/v1/balances") {
			md["x-route-id"] = app.OrderRouteId
			logrusLogger.Debugln("/v1/balances adding order route id", app.OrderRouteId)
		} else if strings.HasPrefix(r.URL.String(), "/v1/assets") {
			md["x-route-id"] = app.OrderRouteId
			logrusLogger.Debugln("/v1/assets adding order route id", app.OrderRouteId)
		} else {
			logrusLogger.Warnf("%s is an unknown route", r.URL.String())
		}

		return metadata.New(md)
	}))

	gwmux.HandlePath("GET", "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		logrusLogger.Debugln("responding to health check")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok\n")
	})

	// Websocket Endpoint
	pool := websocket.NewPool(app)
	go pool.Start()
	logrusLogger.Debugf("created pool and redis client - %v - %v", pool, pool.Redis)
	status := pool.Redis.Ping()
	logrusLogger.Debugf("redis connection status -%v", status)

	gwmux.HandlePath("GET", "/ws", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		serveWs(pool, w, r)
	})

	// Register Service Handlers

	err = profile.RegisterProfileServiceHandler(context.Background(), gwmux, pConn)
	if err != nil {
		logrusLogger.Fatalln("Failed to register profile:", err)
	}

	err = v1.RegisterOrderServiceHandler(context.Background(), gwmux, oConn)
	if err != nil {
		logrusLogger.Fatalln("Failed to register order:", err)
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
		"https://dev.neoworks.xyz",
		"https://api-dev.neoworks.xyz",
		fmt.Sprintf("https://localhost:%s", app.Port),
		fmt.Sprintf("http://localhost:%s", app.Port),
		"http://localhost:4200",
		"https://localhost:4200",
	})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	openCensusHandler := &ochttp.Handler{
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(gwmux),
	}

	logrusLogger.Debugf("starting http - %v - %v - %v", originsOk, headersOk, methodsOk)
	gwServer := &http.Server{
		Handler:      openCensusHandler,
		Addr:         fmt.Sprintf(":%s", app.Port),
		WriteTimeout: 40 * time.Second,
		ReadTimeout:  40 * time.Second,
	}

	logrusLogger.Warnf("checking grpc dials")
	testOrderDial(app)
	testProfileDial(app)

	//assetPriceUpdater(*pool)

	logrusLogger.Warnf("started gRPC-Gateway on - %v", app.Port)

	go func() {
		if app.Env == "local" {
			if err := gwServer.ListenAndServe(); err != nil {
				logrusLogger.Fatalln("ListenAndServe: ", err)
			}
			logrusLogger.Warnf("started http")
		} else {
			if err := gwServer.ListenAndServeTLS("server.crt", "server.key"); err != nil {
				logrusLogger.Fatalln("ListenAndServeTLS: ", err)
			}
			logrusLogger.Warnf("started https")
		}
	}()

	return gwServer, nil
}
