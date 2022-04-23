package types

type Expression struct {
	predicates        []Predicate
	logicalExpression *LogicalExpression
}

type LogicalExpression struct {
	connector  string
	predicates []ExpressionPredicate
}

type ExpressionPredicate struct {
	predicate         Predicate
	logicalExpression *LogicalExpression
}

type Predicate struct {
	Id string
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
