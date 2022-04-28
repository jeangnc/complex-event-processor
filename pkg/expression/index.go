package expression

import (
	"fmt"
	"sort"

	"github.com/jeangnc/complex-event-processor/pkg/tree"
	"github.com/jeangnc/complex-event-processor/pkg/types"

	"golang.org/x/exp/maps"
)

const OPERATOR_EQUAL string = "equal"
const OPERATOR_DIFFERENT string = "not_equal"
const OPERATOR_LESS_THAN string = "less_than"
const OPERATOR_GREATER_THAN string = "greater_than"
const OPERATOR_LESS_THAN_OR_EQUAL = "less_than_or_equal"
const OPERATOR_GREATER_THAN_OR_EQUAL = "greater_than_or_equal"

type Index struct {
	expressionMap map[string][]types.Expression
	predicateTree tree.Node
}

func NewIndex() Index {
	return Index{
		expressionMap: map[string][]types.Expression{},
		predicateTree: tree.NewNode(),
	}
}

func (i Index) SearchImpactedPredicates(e types.Event) types.Impact {
	keys := extractPayloadKeys(e)
	values := i.predicateTree.Values(keys)

	result := map[string]bool{}

	for _, v := range values {
		p := v.(types.Predicate)
		r := evaluateConditions(e, p)

		// immutable false results are irrelevant
		if !p.Immutable || r {
			result[p.Id] = r
		}
	}

	return types.Impact{Predicates: result}
}

func (i Index) FilterImpactedExpressions(c types.Changes) []types.Expression {
	r := make([]types.Expression, 0, 0)

	for p, _ := range c.Predicates {
		es, ok := i.expressionMap[p]

		if !ok {
			continue
		}

		r = append(r, es...)
	}

	return r
}

func (i *Index) Append(e types.Expression) {
	for _, p := range e.LogicalExpression.Predicates() {
		keys := append([]string{e.TenantId}, extractPredicateKeys(p)...)

		n := i.predicateTree.Traverse(keys)
		n.Set(p.Id, *p)

		if _, ok := i.expressionMap[p.Id]; !ok {
			i.expressionMap[p.Id] = make([]types.Expression, 0)
		}
		i.expressionMap[p.Id] = append(i.expressionMap[p.Id], e)
	}
}

func extractPayloadKeys(e types.Event) []string {
	fields := maps.Keys(e.Payload)
	sort.Strings(fields)

	keys := append([]string{e.TenantId, e.Type}, fields...)

	return keys
}

func extractPredicateKeys(p *types.Predicate) []string {
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
			return value == expectedValue
		case OPERATOR_DIFFERENT:
			return value != expectedValue
		case OPERATOR_LESS_THAN:
			// TODO: how to cast?
		case OPERATOR_GREATER_THAN:
			// TODO: how to cast?
		case OPERATOR_LESS_THAN_OR_EQUAL:
			// TODO: how to cast?
		case OPERATOR_GREATER_THAN_OR_EQUAL:
			// TODO: how to cast?
		default:
			panic(fmt.Sprintf("invalid operator %s", c.Operator))
		}
	}

	return true
}
