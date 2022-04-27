package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/mutation"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

var entity = types.Entity{}

func NewEventHandler(index expression.Index) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
