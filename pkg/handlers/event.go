package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/mutation"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

var entity = types.Entity{}

func NewEventHandler(index expression.Index) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var e types.Event

		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			fmt.Fprintf(w, "Invalid event")
		}

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
