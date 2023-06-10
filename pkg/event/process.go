package event

import (
	"context"

	"github.com/jeangnc/complex-event-processor/pkg/expression"
	"github.com/jeangnc/complex-event-processor/pkg/state"
	"github.com/jeangnc/complex-event-processor/pkg/types"
)

func Process(index *expression.Index, repository state.Repository, event types.Event) map[string]bool {
	ctx := context.Background()

	impact := index.SearchImpactedPredicates(event)
	repository.Save(ctx, event, impact)

	expressions := index.FilterImpactedExpressions(impact)
	states, _ := repository.Load(ctx, event, expressions)

	response := map[string]bool{}
	for e, v := range states {
		response[e.Id] = expression.EvaluateExpression(v, e)
	}

	return response
}
