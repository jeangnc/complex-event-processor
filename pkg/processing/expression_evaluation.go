package processing

import (
	"fmt"

	util "github.com/jeangnc/complex-event-processor/pkg/util"
)

const CONNECTOR_AND string = "and"
const CONNECTOR_OR string = "or"

type LogicalExpression struct {
	connector  string
	predicates []ExpressionPredicate
}

type ExpressionPredicate struct {
	predicate         string
	logicalExpression *LogicalExpression
}

func Evaluate(e Entity, ex Expression) bool {
	return evaluateLogicalExpression(e, ex.logicalExpression)
}

func evaluateLogicalExpression(e Entity, l *LogicalExpression) bool {
	values := make([]bool, 0, 0)

	for _, p := range l.predicates {
		if p.logicalExpression != nil {
			values = append(values, evaluateLogicalExpression(e, p.logicalExpression))
			continue
		}

		if value, ok := e.predicates[p.predicate]; !ok {
			values = append(values, false)
		} else {
			values = append(values, value)
		}
	}

	result := false
	switch l.connector {
	case CONNECTOR_AND:
		result = util.SliceAll(values)
	case CONNECTOR_OR:
		result = util.SliceAny(values)
	default:
		// TODO: properly handle this error
		panic(fmt.Sprintf("invalid connector %s", l.connector))
	}

	return result
}
