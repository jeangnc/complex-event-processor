package processing

type LogicalExpression struct {
	connector  string
	predicates []ExpressionPredicate
}

type ExpressionPredicate struct {
	predicate  string
	expression LogicalExpression
}

func Evaluate(e Entity, ex Expression) bool {
	return true
}
