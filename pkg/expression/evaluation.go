package expression

import (
	"fmt"

	"github.com/jeangnc/complex-event-processor/pkg/types"
	util "github.com/jeangnc/complex-event-processor/pkg/util"
)

func EvaluateExpression(e types.Entity, ex types.Expression) bool {
	return evaluateLogicalExpression(e, &ex.LogicalExpression)
}

func evaluateLogicalExpression(e types.Entity, l *types.LogicalExpression) bool {
	values := make([]bool, 0, 0)

	for _, o := range l.Operands {
		if o.LogicalExpression != nil {
			values = append(values, evaluateLogicalExpression(e, o.LogicalExpression))
			continue
		}

		value, ok := e.Predicates[o.Predicate.Id]
		if !ok {
			value = false
		}

		if o.Negated {
			value = !value
		}

		values = append(values, value)
	}

	result := false
	switch l.Connector {
	case types.CONNECTOR_AND:
		result = util.SliceAll(values)
	case types.CONNECTOR_OR:
		result = util.SliceAny(values)
	default:
		panic(fmt.Sprintf("invalid connector %s", l.Connector))
	}

	return result
}
