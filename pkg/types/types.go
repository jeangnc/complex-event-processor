package types

const CONNECTOR_AND string = "and"
const CONNECTOR_OR string = "or"

type Event struct {
	Id        string                 `json:"id"`
	TenantId  string                 `json:"tenant_id"`
	EntityId  string                 `json:"entity_id"`
	Type      string                 `json:"type"`
	Timestamp int64                  `json:"timestamp"`
	Payload   map[string]interface{} `json:"payload"`
}

type State struct {
	Predicates map[string]bool
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

type Notification struct {
	Id           string `json:"id"`
	TenantId     string `json:"tenant_id"`
	EntityId     string `json:"entity_id"`
	ExpressionId string `json:"expression_id"`
	Timestamp    int64  `json:"timestamp"`
	State        bool   `json:"state"`
}

type Expression struct {
	Id                string            `json:"id"`
	TenantId          string            `json:"tenant_id"`
	LogicalExpression LogicalExpression `json:"logical_expression"`
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
	Immutable  bool        `json:"immutable"`
}

type Condition struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}
