package types

const CONNECTOR_AND string = "and"
const CONNECTOR_OR string = "or"

type Event struct {
	Id        string                 `json:"id"`
	Type      string                 `json:"type"`
	Timestamp int64                  `json:"timestamp"`
	Payload   map[string]interface{} `json:"payload"`
}

type State struct {
	Predicates map[string]bool
}

type Impact struct {
	Predicates map[string]bool
}

type Expression struct {
	Id                string            `json:"id"`
	LogicalExpression LogicalExpression `json:"logical_expression"`
	Window            int64             `json:"window"`
}

type Operand struct {
	Negated           bool
	Predicate         *Predicate         `json:"predicate,omitempty"`
	LogicalExpression *LogicalExpression `json:"logical_expression,omitempty"`
}

type Predicate struct {
	Id         string      `json:"id"`
	EventType  string      `json:"event_type"`
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}
