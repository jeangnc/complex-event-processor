package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/handlers"
)

const (
	address = "127.0.0.1:8080"
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

	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)
	router.Use(loggingMiddleware)
	router.HandleFunc("/event", handlers.NewEventHandler(&index)).Methods("POST")
	router.HandleFunc("/expression", handlers.NewExpressionHandler(&index)).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Listening to ", address)
	log.Fatal(srv.ListenAndServe())
}
