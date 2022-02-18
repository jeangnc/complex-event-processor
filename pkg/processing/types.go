package processing

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

type Entity struct {
	predicates map[string]bool
}

type Impact struct {
	predicates map[string]bool
}

type Changes struct {
	predicates map[string]bool
}
