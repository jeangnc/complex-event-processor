package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

func NewExpressionHandler(index *expression.Index) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var e types.Expression

		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			fmt.Fprintf(w, "Invalid event")
		}

		index.Append(e)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
