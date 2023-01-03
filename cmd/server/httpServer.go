package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/coinbase-samples/ib-api-go/auth"
	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/log"
	"github.com/coinbase-samples/ib-api-go/model"
	asset "github.com/coinbase-samples/ib-api-go/pkg/pbs/asset/v1"
	balance "github.com/coinbase-samples/ib-api-go/pkg/pbs/balance/v1"
	order "github.com/coinbase-samples/ib-api-go/pkg/pbs/order/v1"
	profile "github.com/coinbase-samples/ib-api-go/pkg/pbs/profile/v1"
	"github.com/coinbase-samples/ib-api-go/websocket"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func setupHttp(ctx context.Context, app config.AppConfig, aw auth.Middleware) (*http.Server, error) {
	oConn, pConn := setupGrpcDials(ctx, app)
	gwmux := runtime.NewServeMux(makeMetadataOption(app))

	gwmux.HandlePath("GET", "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		log.Debug("responding to health check")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok\n")
	})

	// Websocket Endpoint
	pool := websocket.NewPool(app)
	go pool.Start()
	log.Debugf("created pool and redis client - %v - %v", pool, pool.Redis)
	status := pool.Redis.Ping()
	log.Debugf("redis connection status -%v", status)
	assetPriceUpdater(*pool)

	gwmux.HandlePath("GET", "/ws", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		serveWs(ctx, pool, w, r)
	})

	// Register Service Handlers
	registerServiceHandlers(ctx, gwmux, pConn, oConn)

	gwServer := &http.Server{
		Handler:      makeCorsHandler(ctx, app)(makeContextLogger()(aw.MakeHttpHandler()(gwmux))),
		Addr:         fmt.Sprintf(":%s", app.Port),
		WriteTimeout: 40 * time.Second,
		ReadTimeout:  40 * time.Second,
	}

	log.Debugf("starting gRPC-Gateway on - %v", app.Port)

	go func() {
		if app.Env == "local" {
			if err := gwServer.ListenAndServe(); err != nil {
				log.Fatal("ListenAndServe: ", err)
			}
			log.Debug("started http")
		} else {
			if err := gwServer.ListenAndServeTLS("server.crt", "server.key"); err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
			log.Debug("started https")
		}
	}()

	return gwServer, nil
}

func makeCorsHandler(ctx context.Context, app config.AppConfig) func(http.Handler) http.Handler {
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
	log.Debugf("setting cors headers - %v - %v - %v", originsOk, headersOk, methodsOk)
	return handlers.CORS(originsOk, headersOk, methodsOk)
}

func makeContextLogger() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l := log.NewEntry()
			requestId := uuid.New()
			ctx := log.ToContext(
				context.WithValue(r.Context(), model.RequestCtxKey, requestId),
				l.WithField("requestId", requestId),
			)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func registerServiceHandlers(ctx context.Context, gwmux *runtime.ServeMux, pConn *grpc.ClientConn, oConn *grpc.ClientConn) {
	err := profile.RegisterProfileServiceHandler(context.Background(), gwmux, pConn)
	if err != nil {
		log.Fatal("Failed to register profile:", err)
	}
	err = order.RegisterOrderServiceHandler(context.Background(), gwmux, oConn)
	if err != nil {
		log.Fatal("Failed to register order:", err)
	}
	err = balance.RegisterBalanceServiceHandler(context.Background(), gwmux, oConn)
	if err != nil {
		log.Fatal("Failed to register balance:", err)
	}
	err = asset.RegisterAssetServiceHandler(context.Background(), gwmux, oConn)
	if err != nil {
		log.Fatal("Failed to register asset:", err)
	}
}

func setupGrpcDials(ctx context.Context, app config.AppConfig) (oConn *grpc.ClientConn, pConn *grpc.ClientConn) {
	dialCreds := getGrpcCredentials(app)
	log.Debug("dialing order manager")
	oConn, err := orderConn(app, dialCreds)
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}
	log.Debug("Connected to order manager")

	log.Debug("dialing profile")
	pConn, err = profileConn(app, dialCreds)
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}
	log.Debug("Connected to profile")
	// verify can dial downstream service
	testOrderDial(app)
	testProfileDial(app)

	return oConn, pConn
}

func makeMetadataOption(app config.AppConfig) runtime.ServeMuxOption {
	return runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
		md := make(map[string]string)
		if method, ok := runtime.RPCMethod(ctx); ok {
			md["method"] = method // /grpc.gateway.examples.internal.proto.examplepb.LoginService/Login
		}
		if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
			md["pattern"] = pattern // /v1/example/login
		}

		if strings.HasPrefix(r.URL.String(), "/v1/profile") {
			md["x-route-id"] = app.UserRouteId
			log.Debug("/v1/profile adding profile route id", app.UserRouteId)
		} else if strings.HasPrefix(r.URL.String(), "/v1/order") {
			md["x-route-id"] = app.OrderRouteId
			log.Debug("/v1/order adding order route id", app.OrderRouteId)
		} else if strings.HasPrefix(r.URL.String(), "/v1/balances") {
			md["x-route-id"] = app.OrderRouteId
			log.Debug("/v1/balances adding order route id", app.OrderRouteId)
		} else if strings.HasPrefix(r.URL.String(), "/v1/assets") {
			md["x-route-id"] = app.OrderRouteId
			log.Debug("/v1/assets adding order route id", app.OrderRouteId)
		} else {
			log.Warnf("%s is an unknown route", r.URL.String())
		}

		md["requestId"] = ctx.Value(model.RequestCtxKey).(string)

		return metadata.New(md)
	})
}
