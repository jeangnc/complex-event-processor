package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/handlers"
	"github.com/jeangnc/complex-event-processor/pkg/state"
	muxprom "gitlab.com/msvechla/mux-prometheus/pkg/middleware"
)

const (
	port = ":8080"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		log.Print("Received: ", string(bodyBytes))
		next.ServeHTTP(w, r)
	})
}

func main() {
	index := expression.NewIndex("./tmp/expressions")
	index.Load()

	repository := state.NewRedisRepository("redis:6379", "./lib.lua")

	instrumentation := muxprom.NewDefaultInstrumentation()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)
	router.Use(loggingMiddleware)
	router.Use(instrumentation.Middleware)

	router.HandleFunc("/event", handlers.NewEventHandler(&index, &repository)).Methods("POST")
	router.HandleFunc("/expression", handlers.NewExpressionHandler(&index)).Methods("POST")
	router.Path("/metrics").Handler(promhttp.Handler())

	srv := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Listening to ", port)
	log.Fatal(srv.ListenAndServe())
}
