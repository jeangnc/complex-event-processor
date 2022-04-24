package expression

import (
	"fmt"
	"sort"

	"github.com/jeangnc/complex-event-processor/pkg/tree"
	"github.com/jeangnc/complex-event-processor/pkg/types"

	"golang.org/x/exp/maps"
)

const OPERATOR_EQUAL string = "eq"
const OPERATOR_DIFFERENT string = "not_eq"

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

	r := map[string]bool{}

	for _, v := range values {
		p := v.(types.Predicate)
		r[p.Id] = evaluateConditions(e, p)
	}

	return types.Impact{Predicates: r}
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

func evaluateConditions(e types.Event, p types.Predicate) bool {
	for _, c := range p.Conditions {
		value, ok := e.Payload[c.Field]
		expectedValue := c.Value

		if !ok {
			return false
		}

		switch c.Operator {
		case OPERATOR_EQUAL:
			if value != expectedValue {
				return false
			}
		case OPERATOR_DIFFERENT:
			if value == expectedValue {
				return false
			}
		default:
			// TODO: properly handle this error
			panic(fmt.Sprintf("invalid operator %s", c.Operator))
		}
	}

	return true
}
