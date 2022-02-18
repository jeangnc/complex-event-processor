package processing

const CONNECTOR_AND string = "and"
const CONNECTOR_OR string = "or"

type Expression struct {
	predicates        []string
	logicalExpression *LogicalExpression
}

type LogicalExpression struct {
	connector  string
	predicates []ExpressionPredicate
}

type ExpressionPredicate struct {
	predicate         string
	logicalExpression *LogicalExpression
}

func FilterImpacted(c Changes, es []Expression) []Expression {
	r := make([]Expression, 0, 0)

	for _, e := range es {
		for _, p := range e.predicates {
			if _, ok := c.predicates[p]; ok {
				r = append(r, e)
			}
		}
	}

	return r
}
