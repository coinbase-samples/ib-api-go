package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"time"

	"github.com/coinbase-samples/ib-api-go/config"
	grpcHandlers "github.com/coinbase-samples/ib-api-go/handlers"
	"github.com/coinbase-samples/ib-api-go/pkg/pbs/ledger"
	v1 "github.com/coinbase-samples/ib-api-go/pkg/pbs/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func gRPCListen(app config.AppConfig, aw authMiddleware) {
	tracer := otel.Tracer("ib-api-go")
	logLevel, _ := log.ParseLevel(app.LogLevel)
	logrusLogger.SetLevel(logLevel)
	//setup otel
	tp, err := config.Init(app)
	logrusLogger.Debugln("Started otel", tp)
	if err != nil {
		logrusLogger.Fatal(err)
	}

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logrusLogger.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	/*
		// if local export both grpc and http endpoints
		activePort := app.GrpcPort

		//setup conn

		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", activePort))
		if err != nil {
			logrusLogger.Fatalln("Failed to listen for gRPC: %v", err)
		}
	*/
	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	// See example setup here https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/logging/logrus/examples_test.go
	logrusEntry := log.NewEntry(logrusLogger)
	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	var s *grpc.Server

	if app.Env != "local" {
		// load tls for grpc
		tlsCredentials, err := loadCredentials()
		if err != nil {
			logrusLogger.Fatalln("Cannot load TLS credentials: ", err)
		}

		s = grpc.NewServer(
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				otelgrpc.UnaryServerInterceptor(),
				grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
				aw.InterceptorNew(),
				grpc_validator.UnaryServerInterceptor(),
				grpc_recovery.UnaryServerInterceptor(),
			)),
			grpc.Creds(tlsCredentials),
		)
	} else {

		s = grpc.NewServer(
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				otelgrpc.UnaryServerInterceptor(),
				grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
				aw.InterceptorNew(),
				grpc_validator.UnaryServerInterceptor(),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		)
	}

	v1.RegisterProfileServiceServer(s, &grpcHandlers.ProfileServer{Tracer: tracer})

	//register grpc protos
	oConn, err := orderConn(app)
	if err != nil {
		logrusLogger.Warnln("Could not get order grpc connection")
	}
	client := ledger.NewLedgerClient(oConn)
	v1.RegisterBalanceServiceServer(s, &grpcHandlers.BalanceServer{Tracer: tracer, OrderClient: client})

	//setup health server
	healthServer := health.NewServer()
	healthServer.SetServingStatus("grpc.health.v1.Health", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(s, healthServer)

	reflection.Register(s)

	/*
		fmt.Printf("gRPC Server starting on port %s\n", activePort)
			go func() {
				if err := s.Serve(lis); err != nil {
					fmt.Printf("Failed to listen for gRPC: %v", err)
				}

			}()
	*/
	testOrderDial(app)
	testProfileDial(app)

	gwServer, err := setupHttp(app, s)
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

func loadCredentials() (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(
		&tls.Config{
			Certificates: []tls.Certificate{cert},
			ClientAuth:   tls.NoClientCert,
		},
	), nil
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("ca-cert.pem")
	if err != nil {
		logrusLogger.Warnf("missing cert - %v", err)
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		logrusLogger.Warnf("failed to append cert - %v", err)
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	logrusLogger.Warnf("configured cert - %v", config)
	return credentials.NewTLS(config), nil
}
