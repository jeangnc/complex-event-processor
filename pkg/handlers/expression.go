package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

func NewExpressionHandler(index expression.Index) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Invalid expression")
		}

		var ex types.Expression
		json.Unmarshal(body, &ex)

		index.Append(ex)

		w.WriteHeader(http.StatusCreated)
	}
}
