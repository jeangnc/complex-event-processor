package event

import (
	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/mutation"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

func Process(index *expression.Index, entity types.Entity, event types.Event) (types.Entity, map[string]bool) {
	impacts := index.SearchImpactedPredicates(event)
	newEntity, changes := mutation.Process(entity, impacts)
	exps := index.FilterImpactedExpressions(changes)

	response := map[string]bool{}
	for _, ex := range exps {
		response[ex.Id] = expression.EvaluateExpression(newEntity, ex)
	}
	return newEntity, response
}
