package mutation

import (
	"testing"
)

// Tests whether a set of changes impacted the expression
func TestImpactIdentification(t *testing.T) {
	type testCase struct {
		description    string
		changes        Changes
		expression     Expression
		expectedResult bool
	}

	e := Expression{
		predicates: []string{
			"test",
		},
	}

	testCases := []testCase{
		testCase{
			description: "when the expression was impacted",
			changes: Changes{
				predicates: map[string]bool{
					"test": false,
				},
			},
			expression:     e,
			expectedResult: true,
		},
		testCase{
			description: "when the expression was not impacted",
			changes: Changes{
				predicates: map[string]bool{
					"test2": false,
				},
			},
			expression:     e,
			expectedResult: false,
		},
	}

	for _, s := range testCases {
		t.Run(s.description, func(t *testing.T) {
			result := Impacted(s.changes, s.expression)
			if result != s.expectedResult {
				t.Fatalf(`Failed: %s %v`, s.description, s.expectedResult)
			}
		})

	}

}

// TODO: ensure relations are considered too
