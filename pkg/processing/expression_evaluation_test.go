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

	andEx := Expression{
		logicalExpression: &LogicalExpression{
			connector: CONNECTOR_AND,
			predicates: []ExpressionPredicate{
				ExpressionPredicate{predicate: "a"},
				ExpressionPredicate{predicate: "b"},
			},
		},
	}
	orEx := Expression{
		logicalExpression: &LogicalExpression{
			connector: CONNECTOR_OR,
			predicates: []ExpressionPredicate{
				ExpressionPredicate{predicate: "a"},
				ExpressionPredicate{predicate: "b"},
			},
		},
	}

	testCases := []testCase{
		testCase{
			description: "'AND' expression with truthy result",
			entity: Entity{
				predicates: map[string]bool{
					"a": true,
					"b": true,
				},
			},
			expression:     andEx,
			expectedResult: true,
		},
		testCase{
			description: "'AND' expression with falsey result",
			entity: Entity{
				predicates: map[string]bool{
					"a": true,
				},
			},
			expression:     andEx,
			expectedResult: false,
		},
		testCase{
			description: "'OR' expression with truthy result",
			entity: Entity{
				predicates: map[string]bool{
					"a": true,
					"b": false,
				},
			},
			expression:     orEx,
			expectedResult: true,
		},
		testCase{
			description: "'OR' expression with falsey result",
			entity: Entity{
				predicates: map[string]bool{},
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
		entity         Entity
		expression     Expression
		expectedResult bool
	}

	andEx := Expression{
		logicalExpression: &LogicalExpression{
			connector: CONNECTOR_AND,
			predicates: []ExpressionPredicate{
				ExpressionPredicate{predicate: "a"},
				ExpressionPredicate{
					logicalExpression: &LogicalExpression{
						connector: CONNECTOR_AND,
						predicates: []ExpressionPredicate{
							ExpressionPredicate{predicate: "b"},
						},
					},
				},
			},
		},
	}
	orEx := Expression{
		logicalExpression: &LogicalExpression{
			connector: CONNECTOR_OR,
			predicates: []ExpressionPredicate{
				ExpressionPredicate{predicate: "a"},
				ExpressionPredicate{
					logicalExpression: &LogicalExpression{
						connector: CONNECTOR_AND,
						predicates: []ExpressionPredicate{
							ExpressionPredicate{predicate: "b"},
						},
					},
				},
			},
		},
	}

	testCases := []testCase{
		testCase{
			description: "'AND' expression with truthy result",
			entity: Entity{
				predicates: map[string]bool{
					"a": true,
					"b": true,
				},
			},
			expression:     andEx,
			expectedResult: true,
		},
		testCase{
			description: "'AND' expression with falsey result",
			entity: Entity{
				predicates: map[string]bool{
					"a": true,
				},
			},
			expression:     andEx,
			expectedResult: false,
		},
		testCase{
			description: "'OR' expression with truthy result",
			entity: Entity{
				predicates: map[string]bool{
					"a": true,
					"b": false,
				},
			},
			expression:     orEx,
			expectedResult: true,
		},
		testCase{
			description: "'OR' expression with falsey result",
			entity: Entity{
				predicates: map[string]bool{},
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
