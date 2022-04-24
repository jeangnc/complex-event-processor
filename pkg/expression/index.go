package expression

import (
	"sort"

	"github.com/jeangnc/complex-event-processor/pkg/tree"
	"github.com/jeangnc/complex-event-processor/pkg/types"

	"golang.org/x/exp/maps"
)

type Index struct {
	expressionMap PredicateExpressionMap
	predicateTree tree.Node
}

type PredicateExpressionMap map[string][]*types.Expression

func NewIndex() Index {
	return Index{
		expressionMap: PredicateExpressionMap{},
		predicateTree: tree.NewNode(),
	}
}

func (i *Index) Search(e types.Event) types.Impact {
	keys := extractPayloadKeys(e)
	values := i.predicateTree.Values(keys)

	predicates := map[string]bool{}
	for _, v := range values {
		p := v.(types.Predicate)
		predicates[p.Id] = true
	}

	return types.Impact{Predicates: predicates}
}

func (i *Index) Append(e types.Expression) {
	for _, p := range e.Predicates {
		keys := append([]string{e.TenantId}, extractPredicateKeys(p)...)

		n := i.predicateTree.Traverse(keys)
		n.Set(p.Id, p)

		if _, ok := i.expressionMap[p.Id]; !ok {
			i.expressionMap[p.Id] = make([]*types.Expression, 0)
		}
		i.expressionMap[p.Id] = append(i.expressionMap[p.Id], &e)
	}
}

func extractPayloadKeys(e types.Event) []string {
	fields := maps.Keys(e.Payload)
	sort.Strings(fields)

	keys := append([]string{e.TenantId, e.Type}, fields...)

	return keys
}

func extractPredicateKeys(p types.Predicate) []string {
	fields := make([]string, 0, len(p.Conditions))

	for _, c := range p.Conditions {
		fields = append(fields, c.Field)
	}

	sort.Strings(fields)

	keys := []string{p.EventType}
	keys = append(keys, fields...)

	return keys
}
