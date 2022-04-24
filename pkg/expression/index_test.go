package expression

import (
	"reflect"
	"testing"

	"github.com/jeangnc/complex-event-processor/pkg/types"
)

func TestIndex(t *testing.T) {
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

	result := i.Search(e)

	expectedResult := types.Impact{
		Predicates: map[string]bool{
			"my-predicate": true,
		},
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf(`Failed to search impacted expressions: %v %v`, result, expectedResult)
	}
}
