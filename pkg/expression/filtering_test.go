package expression

import (
	"reflect"
	"testing"
)

// Tests whether a set of changes impacted the expression
func TestImpactIdentification(t *testing.T) {
	type testCase struct {
		description    string
		changes        Changes
		expressions    []Expression
		expectedResult []Expression
	}

	e := Expression{
		predicates: []Predicate{
			Predicate{Id: "test"},
		},
	}
	es := []Expression{e}

	testCases := []testCase{
		testCase{
			description: "when the expression was impacted",
			changes: Changes{
				predicates: map[string]bool{
					"test": false,
				},
			},
			expressions:    es,
			expectedResult: []Expression{e},
		},
		testCase{
			description: "when the expression was not impacted",
			changes: Changes{
				predicates: map[string]bool{
					"test2": false,
				},
			},
			expressions:    es,
			expectedResult: []Expression{},
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
