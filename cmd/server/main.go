package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cfluke-cb/ib-client-api/config"
	restHandlers "github.com/cfluke-cb/ib-client-api/internal/handlers"
	"github.com/cfluke-cb/ib-client-api/model"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func setupRoutes(router *mux.Router, app config.AppConfig) {

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok\n")
	})

	router.HandleFunc("/v1/profile/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		body, err := restHandlers.FetchProfile(id)

		if err != nil {
			fmt.Fprintf(w, "%+v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "")
			return
		}

		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(body)
		io.WriteString(w, string(response))
	})

	router.HandleFunc("/v1/order", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Could not parse request body")
			return
		}
		var order = &model.OrderRequest{}
		json.Unmarshal(b, order)
		body, err := restHandlers.PlaceOrder(*order)

		if err != nil {
			fmt.Fprintf(w, "%+v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "")
			return
		}

		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(body)
		io.WriteString(w, string(response))
	})
}

func main() {

	var wait time.Duration
	var app config.AppConfig

	config.Setup(&app)

	router := mux.NewRouter()
	setupRoutes(router, app)
	port := "8443"
	if app.Port != "" {
		port = app.Port
	}
	fmt.Printf("starting listener on: %s\n", port)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"https://api.neoworks.xyz",
		"https://api-dev.neoworks.xyz",
		fmt.Sprintf("https://localhost:%s", port),
		fmt.Sprintf("http://localhost:%s", port),
		"https://localhost:4200",
		"http://localhost:4200"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	srv := &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(router),
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 40 * time.Second,
		ReadTimeout:  40 * time.Second,
	}

	go func() {
		if app.Env != "local" {
			if err := srv.ListenAndServeTLS("server.crt", "server.key"); err != nil {

				log.Fatal("ListenAndServeTLS: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil {

				log.Fatal("ListenAndServe: ", err)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("stopping")
	os.Exit(0)
}
