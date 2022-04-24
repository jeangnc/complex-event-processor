package expression

import (
	"reflect"
	"testing"

	"github.com/jeangnc/complex-event-processor/pkg/types"
)

// Test wheter we can identify which predicates were impacted
func TestImpactedPredicatesSearch(t *testing.T) {
	tenantId := "1"

	p := types.Predicate{
		Id:        "my-predicate",
		EventType: "EMAIL_OPENED",
		Conditions: []types.Condition{
			types.Condition{
				Field:    "email",
				Operator: OPERATOR_EQUAL,
				Value:    "test",
			},
		},
	}

	ex := types.Expression{
		TenantId:   tenantId,
		Predicates: []types.Predicate{p},
		LogicalExpression: &types.LogicalExpression{
			Connector: CONNECTOR_AND,
			Predicates: []types.ExpressionPredicate{
				types.ExpressionPredicate{Predicate: p},
			},
		},
	}

	e := types.Event{
		TenantId: tenantId,
		Type:     "EMAIL_OPENED",
		Payload: map[string]interface{}{
			"email": "test",
		},
	}

	i := NewIndex()
	i.Append(ex)

	result := i.SearchImpactedPredicates(e)

	expectedResult := types.Impact{
		Predicates: map[string]bool{
			"my-predicate": true,
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf(`Failed to search impacted expressions: %v %v`, result, expectedResult)
	}
}

// Tests whether a set of changes impacted the expression
func TestImpactedExpressionsFilter(t *testing.T) {
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
			i := NewIndex()

			for _, ex := range s.expressions {
				i.Append(ex)
			}

			result := i.FilterImpactedExpressions(s.changes)
			if !reflect.DeepEqual(result, s.expectedResult) {
				t.Fatalf(`Failed: %s %v %v`, s.description, result, s.expectedResult)
			}
		})
	}
}
