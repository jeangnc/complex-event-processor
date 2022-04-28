package expression

import (
	"reflect"
	"testing"

	"github.com/jeangnc/complex-event-processor/pkg/types"
)

func TestEvaluation(t *testing.T) {
	type testCase struct {
		description    string
		entity         types.Entity
		expression     types.Expression
		expectedResult bool
	}

	andEx := types.Expression{
		LogicalExpression: types.LogicalExpression{
			Connector: types.CONNECTOR_AND,
			Operands: []types.Operand{
				types.Operand{Predicate: &types.Predicate{Id: "a"}},
				types.Operand{Predicate: &types.Predicate{Id: "b"}},
			},
		},
	}
	orEx := types.Expression{
		LogicalExpression: types.LogicalExpression{
			Connector: types.CONNECTOR_OR,
			Operands: []types.Operand{
				types.Operand{Predicate: &types.Predicate{Id: "a"}},
				types.Operand{Predicate: &types.Predicate{Id: "b"}},
			},
		},
	}

	testCases := []testCase{
		testCase{
			description: "'AND' expression with truthy result",
			entity: types.Entity{
				Predicates: map[string]bool{
					"a": true,
					"b": true,
				},
			},
			expression:     andEx,
			expectedResult: true,
		},
		testCase{
			description: "'AND' expression with falsey result",
			entity: types.Entity{
				Predicates: map[string]bool{
					"a": true,
				},
			},
			expression:     andEx,
			expectedResult: false,
		},
		testCase{
			description: "negated 'AND' expression with truthy result",
			entity: types.Entity{
				Predicates: map[string]bool{
					"a": true,
				},
			},
			expression: types.Expression{
				LogicalExpression: types.LogicalExpression{
					Connector: types.CONNECTOR_AND,
					Operands: []types.Operand{
						types.Operand{Predicate: &types.Predicate{Id: "a"}},
						types.Operand{
							Negated:   true,
							Predicate: &types.Predicate{Id: "b"},
						},
					},
				},
			},
			expectedResult: true,
		},
		testCase{
			description: "'OR' expression with truthy result",
			entity: types.Entity{
				Predicates: map[string]bool{
					"a": true,
					"b": false,
				},
			},
			expression:     orEx,
			expectedResult: true,
		},
		testCase{
			description: "'OR' expression with falsey result",
			entity: types.Entity{
				Predicates: map[string]bool{},
			},
			expression:     orEx,
			expectedResult: false,
		},
	}

	for _, s := range testCases {
		t.Run(s.description, func(t *testing.T) {
			result := EvaluateExpression(s.entity, s.expression)
			if !reflect.DeepEqual(result, s.expectedResult) {
				t.Fatalf(`Failed: %s`, s.description)
			}
		})
	}
}

func TestExpressionNesting(t *testing.T) {
	type testCase struct {
		description    string
		entity         types.Entity
		expression     types.Expression
		expectedResult bool
	}

	andEx := types.Expression{
		LogicalExpression: types.LogicalExpression{
			Connector: types.CONNECTOR_AND,
			Operands: []types.Operand{
				types.Operand{Predicate: &types.Predicate{Id: "a"}},
				types.Operand{
					LogicalExpression: &types.LogicalExpression{
						Connector: types.CONNECTOR_AND,
						Operands: []types.Operand{
							types.Operand{Predicate: &types.Predicate{Id: "b"}},
						},
					},
				},
			},
		},
	}
	orEx := types.Expression{
		LogicalExpression: types.LogicalExpression{
			Connector: types.CONNECTOR_OR,
			Operands: []types.Operand{
				types.Operand{Predicate: &types.Predicate{Id: "a"}},
				types.Operand{
					LogicalExpression: &types.LogicalExpression{
						Connector: types.CONNECTOR_AND,
						Operands: []types.Operand{
							types.Operand{Predicate: &types.Predicate{Id: "b"}},
						},
					},
				},
			},
		},
	}

	testCases := []testCase{
		testCase{
			description: "'AND' expression with truthy result",
			entity: types.Entity{
				Predicates: map[string]bool{
					"a": true,
					"b": true,
				},
			},
			expression:     andEx,
			expectedResult: true,
		},
		testCase{
			description: "'AND' expression with falsey result",
			entity: types.Entity{
				Predicates: map[string]bool{
					"a": true,
				},
			},
			expression:     andEx,
			expectedResult: false,
		},
		testCase{
			description: "'OR' expression with truthy result",
			entity: types.Entity{
				Predicates: map[string]bool{
					"a": true,
					"b": false,
				},
			},
			expression:     orEx,
			expectedResult: true,
		},
		testCase{
			description: "'OR' expression with falsey result",
			entity: types.Entity{
				Predicates: map[string]bool{},
			},
			expression:     orEx,
			expectedResult: false,
		},
	}

	for _, s := range testCases {
		t.Run(s.description, func(t *testing.T) {
			result := EvaluateExpression(s.entity, s.expression)
			if !reflect.DeepEqual(result, s.expectedResult) {
				t.Fatalf(`Failed: %s`, s.description)
			}
		})
	}
}
