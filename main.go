package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/handlers"
)

const (
	port = ":8080"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		log.Print("Received: ", string(bodyBytes))
		next.ServeHTTP(w, r)
	})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	index := expression.NewIndex()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	router.Use(jsonMiddleware)
	router.HandleFunc("/event", handlers.NewEventHandler(index)).Methods("POST")
	router.HandleFunc("/expression", handlers.NewExpressionHandler(index)).Methods("POST")

	log.Print("Listening to ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
