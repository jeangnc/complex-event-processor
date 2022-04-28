package main

import (
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
		body, _ := ioutil.ReadAll(r.Body)
		log.Print("Received: ", string(body))
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
