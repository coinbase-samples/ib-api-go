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

	"github.com/coinbase-samples/ib-api-go/config"
	v1 "github.com/coinbase-samples/ib-api-go/pkg/pbs/v1"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func getOrderConnAddress(app config.AppConfig) string {
	if app.IsLocalEnv() {
		return fmt.Sprintf("%s:%s", app.NetworkName, app.OrderGrpcPort)
	}
	return fmt.Sprintf("%s:443", app.NetworkName)
}

func getProfileConnAddress(app config.AppConfig) string {
	if app.IsLocalEnv() {
		return fmt.Sprintf("%s:%s", "0.0.0.0", app.GrpcPort)
	}
	return fmt.Sprintf("%s:%s", "api-internal-dev.neoworks.xyz", app.GrpcPort)
}

func testOrderDial(app config.AppConfig) {
	dialOrderConn := getOrderConnAddress(app)
	grpc.EnableTracing = true
	var clientCreds credentials.TransportCredentials
	if app.Env != "local" {
		clientCreds, _ = loadTLSCredentials()
	} else {
		clientCreds = insecure.NewCredentials()
	}
	conn, err := grpc.Dial(dialOrderConn, grpc.WithTransportCredentials(clientCreds)) //insecure.NewCredentials()))
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
	logrusLogger.Warnf("Greeting: %s", r.Data)
}

func orderConn(app config.AppConfig) (*grpc.ClientConn, error) {

	dialOrderConn := getOrderConnAddress(app)

	var opts []grpc.DialOption
	if app.IsLocalEnv() {
		opts = append(opts, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				InsecureSkipVerify: true,
			}),
		))
	}

	md := metadata.New(map[string]string{"x-route-id": app.OrderRouteId})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	logrusLogger.Warnln("order dial", dialOrderConn, md)
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests

	conn, err := grpc.DialContext(
		ctx,
		dialOrderConn,
		opts...,
	)
	return conn, err
}

func testProfileDial(app config.AppConfig) {
	dialProfileConn := getProfileConnAddress(app)
	grpc.EnableTracing = true

	conn, err := grpc.Dial(dialProfileConn, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := v1.NewProfileServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ReadProfile(ctx, &v1.ReadProfileRequest{Id: "c5af3271-7185-4a52-9d0c-1c4b418317d8"})
	grpc.EnableTracing = false

	if err != nil {
		logrusLogger.Warnf("could not greet profile: %v", err)
		return
	}
	logrusLogger.Warnf("Greeting: %s", r.UserName)
}

func profileConn(app config.AppConfig) (*grpc.ClientConn, error) {
	dialProfileConn := getProfileConnAddress(app)
	logrusLogger.Warnln("connecting to profile localhost grpc", dialProfileConn)
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		dialProfileConn,
		//grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return conn, err
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func setupHttp(app config.AppConfig, grpcServer *grpc.Server) (*http.Server, error) {
	logrusLogger.Warnln("dialing order manager")
	oConn, err := orderConn(app)
	if err != nil {
		logrusLogger.Fatalln("Failed to dial server:", err)
	}
	logrusLogger.Warnln("Connected to order manager")

	logrusLogger.Warnln("dialing profile")
	pConn, err := profileConn(app)
	if err != nil {
		logrusLogger.Fatalln("Failed to dial server:", err)
	}
	logrusLogger.Warnln("Connected to profile")

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

	gwmux.HandlePath("GET", "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		logrusLogger.Warnln("responding to health check")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok\n")
	})

	// Register Service Handlers
	err = v1.RegisterProfileServiceHandler(context.Background(), gwmux, pConn)
	//dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err = v1.RegisterProfileServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8443", dopts)
	if err != nil {
		logrusLogger.Fatalln("Failed to register profile:", err)
	}

	//add route path
	md := metadata.New(map[string]string{"x-route-id": app.OrderRouteId})
	orderHandlerContext := metadata.NewOutgoingContext(context.Background(), md)
	err = v1.RegisterOrderServiceHandler(orderHandlerContext, gwmux, oConn)
	if err != nil {
		logrusLogger.Fatalln("Failed to register order:", err)
	}
	err = v1.RegisterBalanceServiceHandler(orderHandlerContext, gwmux, oConn)
	if err != nil {
		logrusLogger.Fatalln("Failed to register balance:", err)
	}
	err = v1.RegisterAssetServiceHandler(orderHandlerContext, gwmux, oConn)
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

	logrusLogger.Warnf("starting http - %v - %v - %v", originsOk, headersOk, methodsOk)
	gwServer := &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(gwmux),
		Addr:         fmt.Sprintf(":%s", app.Port),
		WriteTimeout: 40 * time.Second,
		ReadTimeout:  40 * time.Second,
	}

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
