package handlers

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

func save(filename string, e []types.Expression) {
	f, _ := os.Create(filename)
	c := gob.NewEncoder(f)
	c.Encode(e)
}

func NewExpressionHandler(index *expression.Index) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var e types.Expression

		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			fmt.Fprintf(w, "Invalid event")
		}

		index.Append(e)

		// TODO: improve this
		save("expressions", index.Expressions())

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
