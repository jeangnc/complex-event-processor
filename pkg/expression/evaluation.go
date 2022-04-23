package expression

import (
	"fmt"

	"github.com/jeangnc/complex-event-processor/pkg/types"
	util "github.com/jeangnc/complex-event-processor/pkg/util"
)

func EvaluateExpression(e types.Entity, ex types.Expression) bool {
	return evaluateLogicalExpression(e, ex.LogicalExpression)
}

func evaluateLogicalExpression(e types.Entity, l *types.LogicalExpression) bool {
	values := make([]bool, 0, 0)

	for _, p := range l.Predicates {
		if p.LogicalExpression != nil {
			values = append(values, evaluateLogicalExpression(e, p.LogicalExpression))
			continue
		}

		value, ok := e.Predicates[p.Predicate.Id]
		if !ok {
			value = false
		}
		values = append(values, value)
	}

	result := false
	switch l.Connector {
	case CONNECTOR_AND:
		result = util.SliceAll(values)
	case CONNECTOR_OR:
		result = util.SliceAny(values)
	default:
		// TODO: properly handle this error
		panic(fmt.Sprintf("invalid connector %s", l.Connector))
	}

	return result
}
