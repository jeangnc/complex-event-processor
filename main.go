package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/handlers"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

const (
	port = ":8080"
)

func load(filename string, e *[]types.Expression) {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	d := gob.NewDecoder(f)
	err = d.Decode(&e)

	if err != nil {
		fmt.Println("Decode error:", err)
	}
}

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
	index := expression.NewIndex()

	es := make([]types.Expression, 0)
	load("expressions", &es)
	for _, e := range es {
		index.Append(e)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)
	router.Use(loggingMiddleware)
	router.HandleFunc("/event", handlers.NewEventHandler(&index)).Methods("POST")
	router.HandleFunc("/expression", handlers.NewExpressionHandler(&index)).Methods("POST")

	log.Print("Listening to ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
