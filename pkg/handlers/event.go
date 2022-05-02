package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeangnc/complex-event-processor/pkg/event"
	"github.com/jeangnc/complex-event-processor/pkg/expression"
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

		newEntity, response := event.Process(index, entity, e)
		entity = newEntity

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
