package main

import (
	"fmt"

	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/mutation"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

const (
	port = ":8080"
)

func main() {
	// send notifications (TODO)

	tenantId := "1"

	p := types.Predicate{
		Id:        "my-predicate",
		EventType: "EMAIL_OPENED",
		Immutable: true,
		Conditions: []types.Condition{
			types.Condition{
				Field:    "email",
				Operator: expression.OPERATOR_EQUAL,
				Value:    "test",
			},
		},
	}

	ex := types.Expression{
		TenantId:   tenantId,
		Predicates: []types.Predicate{p},
		LogicalExpression: types.LogicalExpression{
			Connector: expression.CONNECTOR_AND,
			Predicates: []types.ExpressionPredicate{
				types.ExpressionPredicate{Predicate: p},
			},
		},
	}

	i := expression.NewIndex()
	i.Append(ex)

	entity := types.Entity{}

	e1 := types.Event{
		TenantId: tenantId,
		Type:     "EMAIL_OPENED",
		Payload: map[string]interface{}{
			"email": "test",
		},
	}

	e2 := types.Event{
		TenantId: tenantId,
		Type:     "EMAIL_OPENED",
		Payload: map[string]interface{}{
			"email": "test 2",
		},
	}

	events := []types.Event{e1, e2}

	for _, e := range events {
		impact := i.SearchImpactedPredicates(e)
		newEntity, changes := mutation.Process(entity, impact)
		entity = newEntity
		expressions := i.FilterImpactedExpressions(changes)

		for _, ex := range expressions {
			fmt.Println("Result:", expression.EvaluateExpression(entity, ex))
		}
	}
}
