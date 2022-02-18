package processing

import (
	"reflect"
	"testing"
)

func TestEvaluation(t *testing.T) {
	type testCase struct {
		description    string
		entity         Entity
		expression     Expression
		expectedResult bool
	}

	ex := Expression{
		logicalExpression: LogicalExpression{
			connector: CONNECTOR_AND,
			predicates: []ExpressionPredicate{
				ExpressionPredicate{
					predicate: "test",
				},
			},
		},
		predicates: []string{
			"test",
		},
	}

	testCases := []testCase{
		testCase{
			description: "truthy result",
			entity: Entity{
				predicates: map[string]bool{
					"test": true,
				},
			},
			expression:     ex,
			expectedResult: true,
		},
		testCase{
			description: "falsey result",
			entity: Entity{
				predicates: map[string]bool{},
			},
			expression:     ex,
			expectedResult: false,
		},
	}

	for _, s := range testCases {
		t.Run(s.description, func(t *testing.T) {
			result := Evaluate(s.entity, s.expression)
			if !reflect.DeepEqual(result, s.expectedResult) {
				t.Fatalf(`Failed: %s`, s.description)
			}
		})
	}
}
