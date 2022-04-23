package expression

import (
	"reflect"
	"testing"

	"github.com/jeangnc/complex-event-processor/pkg/types"
)

// Tests whether a set of changes impacted the expression
func TestImpactIdentification(t *testing.T) {
	type testCase struct {
		description    string
		changes        types.Changes
		expressions    []types.Expression
		expectedResult []types.Expression
	}

	e := types.Expression{
		Predicates: []types.Predicate{
			types.Predicate{Id: "test"},
		},
	}
	es := []types.Expression{e}

	testCases := []testCase{
		testCase{
			description: "when the expression was impacted",
			changes: types.Changes{
				Predicates: map[string]bool{
					"test": false,
				},
			},
			expressions:    es,
			expectedResult: []types.Expression{e},
		},
		testCase{
			description: "when the expression was not impacted",
			changes: types.Changes{
				Predicates: map[string]bool{
					"test2": false,
				},
			},
			expressions:    es,
			expectedResult: []types.Expression{},
		},
	}

	for _, s := range testCases {
		t.Run(s.description, func(t *testing.T) {
			result := FilterImpacted(s.changes, s.expressions)
			if !reflect.DeepEqual(result, s.expectedResult) {
				t.Fatalf(`Failed: %s %v`, s.description, s.expectedResult)
			}
		})
	}
}
