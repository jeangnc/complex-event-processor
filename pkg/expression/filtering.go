package expression

import "github.com/jeangnc/complex-event-processor/pkg/types"

func FilterImpacted(c types.Changes, es []types.Expression) []types.Expression {
	r := make([]types.Expression, 0, 0)

	for _, e := range es {
		for _, p := range e.Predicates {
			if _, ok := c.Predicates[p.Id]; ok {
				r = append(r, e)
			}
		}
	}

	return r
}
