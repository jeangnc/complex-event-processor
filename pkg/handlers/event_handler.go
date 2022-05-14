package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeangnc/complex-event-processor/pkg/event"
	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

var entities = map[string]types.Entity{}

func NewEventHandler(index *expression.Index) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var e types.Event

		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid event: ", err)
			return
		}

		entity, ok := entities[e.EntityId]
		if !ok {
			entity = types.Entity{}
		}

		newEntity, response := event.Process(index, entity, e)
		entities[e.EntityId] = newEntity

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
