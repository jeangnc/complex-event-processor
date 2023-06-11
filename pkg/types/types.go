package types

const CONNECTOR_AND string = "and"
const CONNECTOR_OR string = "or"
const CONNECTOR_SEQUENCE string = "sequence"

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
	Within            string            `json:"within"`
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
