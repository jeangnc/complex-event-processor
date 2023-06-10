package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeangnc/complex-event-processor/pkg/event"
	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/state"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

func NewEventHandler(index *expression.Index, repository state.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var e types.Event

		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid event: ", err)
			return
		}

		response := event.Process(index, repository, e)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
