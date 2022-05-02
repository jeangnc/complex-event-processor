package event

import (
	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/mutation"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

func Process(index expression.Index, entity types.Entity, event types.Event) map[string]bool {
	i := index.SearchImpactedPredicates(event)
	entity, c := mutation.Process(entity, i)
	exps := index.FilterImpactedExpressions(c)

	response := map[string]bool{}
	for _, ex := range exps {
		response[ex.Id] = expression.EvaluateExpression(entity, ex)
	}
	return response
}
