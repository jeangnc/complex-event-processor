package types

type Expression struct {
	Predicates        []Predicate
	LogicalExpression *LogicalExpression
}

type LogicalExpression struct {
	Connector  string
	Predicates []ExpressionPredicate
}

type ExpressionPredicate struct {
	Predicate         Predicate
	LogicalExpression *LogicalExpression
}

type Predicate struct {
	Id string
}

type Entity struct {
	Predicates map[string]bool
}

type Impact struct {
	Predicates map[string]bool
}

type Changes struct {
	Predicates map[string]bool
}
