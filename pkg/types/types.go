package types

type Expression struct {
	Id                string            `json:"id"`
	TenantId          string            `json:"tenant_id"`
	Predicates        []Predicate       `json:"predicates"`
	LogicalExpression LogicalExpression `json:"logical_expression"`
}

type Predicate struct {
	Id         string      `json:"id"`
	EventType  string      `json:"event_type"`
	Conditions []Condition `json:"conditions"`
	Immutable  bool        `json:"immutable"`
}

type Condition struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type LogicalExpression struct {
	Connector  string                `json:"connector"`
	Predicates []ExpressionPredicate `json:"predicates"`
}

type ExpressionPredicate struct {
	Predicate         Predicate          `json:"predicate,omitempty"`
	LogicalExpression *LogicalExpression `json:"logical_expression,omitempty"`
}

type Event struct {
	Id       string                 `json:"id"`
	TenantId string                 `json:"tenant_id"`
	Type     string                 `json:"type"`
	Payload  map[string]interface{} `json:"payload"`
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
