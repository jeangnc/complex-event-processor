package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/mutation"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

const (
	port = ":8080"
)

var index = expression.NewIndex()
var entity = types.Entity{}

func expressionHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid expression")
	}

	log.Print("Received an expression: ", string(body))

	var ex types.Expression
	json.Unmarshal(body, &ex)

	index.Append(ex)

	w.WriteHeader(http.StatusCreated)
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid event")
	}

	log.Print("Received an event: ", string(body))

	var e types.Event
	json.Unmarshal(body, &e)

	i := index.SearchImpactedPredicates(e)
	entity, c := mutation.Process(entity, i)
	exps := index.FilterImpactedExpressions(c)

	response := map[string]bool{}
	for _, ex := range exps {
		response[ex.Id] = expression.EvaluateExpression(entity, ex)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/event", eventHandler)
	router.HandleFunc("/expression", expressionHandler)
	log.Print("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
