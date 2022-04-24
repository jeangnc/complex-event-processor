package types

type Expression struct {
	TenantId          string
	Predicates        []Predicate
	LogicalExpression *LogicalExpression
}

type Predicate struct {
	Id         string
	EventType  string
	Conditions []Condition
}

type Condition struct {
	Field    string
	Operator string
	Value    interface{}
}

type LogicalExpression struct {
	Connector  string
	Predicates []ExpressionPredicate
}

type ExpressionPredicate struct {
	Predicate         Predicate
	LogicalExpression *LogicalExpression
}

type Event struct {
	TenantId string
	Type     string
	Payload  map[string]interface{}
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
