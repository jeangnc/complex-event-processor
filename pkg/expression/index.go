package expression

import (
	"encoding/gob"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/jeangnc/complex-event-processor/pkg/tree"
	"github.com/jeangnc/complex-event-processor/pkg/types"
	"github.com/jeangnc/complex-event-processor/pkg/types/value"

	"golang.org/x/exp/maps"
)

const OPERATOR_EQUAL string = "equal"
const OPERATOR_DIFFERENT string = "not_equal"
const OPERATOR_LESS_THAN string = "less_than"
const OPERATOR_GREATER_THAN string = "greater_than"
const OPERATOR_LESS_THAN_OR_EQUAL = "less_than_or_equal"
const OPERATOR_GREATER_THAN_OR_EQUAL = "greater_than_or_equal"

type Index struct {
	filename             string
	expressions          map[string]*types.Expression
	predicateExpressions map[string][]*types.Expression
	predicateTree        tree.Node
	mutex                sync.Mutex
}

func NewIndex(filename string) Index {
	return Index{
		filename:             filename,
		expressions:          map[string]*types.Expression{},
		predicateExpressions: map[string][]*types.Expression{},
		predicateTree:        tree.NewNode(),
	}
}

func NewTemporaryIndex() Index {
	return NewIndex("")
}

func (i *Index) Load() {
	f, err := os.Open(i.filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	var es []types.Expression

	d := gob.NewDecoder(f)
	err = d.Decode(&es)
	if err != nil {
		fmt.Println("Decode error:", err)
	}

	for _, e := range es {
		i.Append(e)
	}
}

func (i Index) Save() {
	e := i.Expressions()
	f, _ := os.Create(i.filename)
	c := gob.NewEncoder(f)
	c.Encode(e)
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
		es, ok := i.predicateExpressions[p]

		if !ok {
			continue
		}

		for _, e := range es {
			r = append(r, *e)
		}
	}

	return r
}

func (i *Index) Append(e types.Expression) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.expressions[e.Id] = &e

	for _, p := range e.LogicalExpression.Predicates() {
		keys := append([]string{e.TenantId}, extractPredicateKeys(p)...)

		n := i.predicateTree.Traverse(keys)
		n.Set(p.Id, *p)

		if _, ok := i.predicateExpressions[p.Id]; !ok {
			i.predicateExpressions[p.Id] = make([]*types.Expression, 0)
		}
		i.predicateExpressions[p.Id] = append(i.predicateExpressions[p.Id], &e)
	}
}

func (i Index) Expressions() []types.Expression {
	r := make([]types.Expression, 0)

	for _, e := range i.expressions {
		r = append(r, *e)
	}

	return r
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
		payloadValue, ok := e.Payload[c.Field]

		if !ok {
			return false
		}

		genericValue := value.NewValue(payloadValue)
		expectedValue := value.NewValue(c.Value)

		switch c.Operator {
		case OPERATOR_EQUAL:
			return genericValue.Equal(expectedValue)
		case OPERATOR_DIFFERENT:
			return genericValue.Different(expectedValue)
		case OPERATOR_LESS_THAN:
			return genericValue.LessThan(expectedValue)
		case OPERATOR_GREATER_THAN:
			return genericValue.GreaterThan(expectedValue)
		case OPERATOR_LESS_THAN_OR_EQUAL:
			return genericValue.LessThanEqual(expectedValue)
		case OPERATOR_GREATER_THAN_OR_EQUAL:
			return genericValue.GreaterThanEqual(expectedValue)
		default:
			panic(fmt.Sprintf("invalid operator %s", c.Operator))
		}
	}

	return true
}
