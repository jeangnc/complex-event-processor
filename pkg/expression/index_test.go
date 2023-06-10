package expression

import (
	"reflect"
	"testing"

	"github.com/jeangnc/complex-event-processor/pkg/types"
)

// Test wheter we can identify which predicates were impacted
func TestImpactedPredicatesSearch(t *testing.T) {
	type testCase struct {
		description    string
		event          types.Event
		predicate      types.Predicate
		expectedResult types.Impact
	}

	event := types.Event{
		Type: "MY_TYPE",
		Payload: map[string]interface{}{
			"string_field": "my-value",
			"int_field":    1,
			"float_field":  1.0,
		},
	}

	testCases := []testCase{
		testCase{
			description: "truthy single condition",
			event:       event,
			predicate: types.Predicate{
				Id:        "my-predicate",
				EventType: "MY_TYPE",
				Conditions: []types.Condition{
					types.Condition{
						Field:    "string_field",
						Operator: OPERATOR_EQUAL,
						Value:    "my-value",
					},
				},
			},
			expectedResult: types.Impact{
				Predicates: map[string]bool{
					"my-predicate": true,
				},
			},
		},
		testCase{
			description: "falsey single condition",
			event:       event,
			predicate: types.Predicate{
				Id:        "my-predicate",
				EventType: "MY_TYPE",
				Conditions: []types.Condition{
					types.Condition{
						Field:    "string_field",
						Operator: OPERATOR_EQUAL,
						Value:    "different-value",
					},
				},
			},
			expectedResult: types.Impact{
				Predicates: map[string]bool{
					"my-predicate": false,
				},
			},
		},
		testCase{
			description: "falsey immutable condition",
			event:       event,
			predicate: types.Predicate{
				Id:        "my-predicate",
				EventType: "MY_TYPE",
				Immutable: true,
				Conditions: []types.Condition{
					types.Condition{
						Field:    "string_field",
						Operator: OPERATOR_EQUAL,
						Value:    "different-test",
					},
				},
			},
			expectedResult: types.Impact{
				Predicates: map[string]bool{},
			},
		},
		testCase{
			description: "empty predicate",
			event:       event,
			predicate: types.Predicate{
				Id:         "my-predicate",
				EventType:  "MY_TYPE",
				Conditions: []types.Condition{},
			},
			expectedResult: types.Impact{
				Predicates: map[string]bool{
					"my-predicate": true,
				},
			},
		},
		testCase{
			description: "different event type",
			event:       event,
			predicate: types.Predicate{
				Id:         "my-predicate",
				EventType:  "MY_OTHER_TYPE",
				Conditions: []types.Condition{},
			},
			expectedResult: types.Impact{
				Predicates: map[string]bool{},
			},
		},
	}

	for _, s := range testCases {
		t.Run(s.description, func(t *testing.T) {
			ex := types.Expression{
				LogicalExpression: types.LogicalExpression{
					Connector: types.CONNECTOR_AND,
					Operands: []types.Operand{
						types.Operand{Predicate: &s.predicate},
					},
				},
			}

			i := NewTemporaryIndex()
			i.Append(ex)

			result := i.SearchImpactedPredicates(s.event)

			if !reflect.DeepEqual(result, s.expectedResult) {
				t.Fatalf(`Failed to search impacted expressions: %v %v`, result, s.expectedResult)
			}
		})
	}
}

// Tests whether a set of changes impacted the expression
func TestImpactedExpressionsFilter(t *testing.T) {
	type testCase struct {
		description    string
		impact         types.Impact
		expressions    []types.Expression
		expectedResult []*types.Expression
	}

	e := types.Expression{
		LogicalExpression: types.LogicalExpression{
			Connector: types.CONNECTOR_AND,
			Operands: []types.Operand{
				types.Operand{Predicate: &types.Predicate{Id: "test"}},
			},
		},
	}

	es := []types.Expression{e}

	testCases := []testCase{
		testCase{
			description: "when the expression was impacted",
			impact: types.Impact{
				Predicates: map[string]bool{
					"test": false,
				},
			},
			expressions:    es,
			expectedResult: []*types.Expression{&e},
		},
		testCase{
			description: "when the expression was not impacted",
			impact: types.Impact{
				Predicates: map[string]bool{
					"test2": false,
				},
			},
			expressions:    es,
			expectedResult: []*types.Expression{},
		},
	}

	for _, s := range testCases {
		t.Run(s.description, func(t *testing.T) {
			i := NewTemporaryIndex()

			for _, ex := range s.expressions {
				i.Append(ex)
			}

			result := i.FilterImpactedExpressions(s.impact)
			if !reflect.DeepEqual(result, s.expectedResult) {
				t.Fatalf(`Failed: %s %v %v`, s.description, result, s.expectedResult)
			}
		})
	}
}
