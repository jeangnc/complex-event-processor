package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/handlers"
)

const (
	port = ":8080"
)

var index = expression.NewIndex()

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/event", handlers.NewEventHandler(index)).Methods("POST")
	router.HandleFunc("/expression", handlers.NewExpressionHandler(index)).Methods("POST")
	log.Print("Listening to ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
