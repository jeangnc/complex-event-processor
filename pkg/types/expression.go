package types

const CONNECTOR_AND string = "and"
const CONNECTOR_OR string = "or"

type Expression struct {
	Id                string            `json:"id"`
	TenantId          string            `json:"tenant_id"`
	LogicalExpression LogicalExpression `json:"logical_expression"`
}

type LogicalExpression struct {
	Connector string    `json:"connector"`
	Operands  []Operand `json:"operands"`
}

type Operand struct {
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

func (le LogicalExpression) Predicates() []*Predicate {
	r := make([]*Predicate, 0)

	for _, o := range le.Operands {
		if o.Predicate != nil {
			r = append(r, o.Predicate)
			continue
		}

		r = append(r, o.LogicalExpression.Predicates()...)
	}

	return r
}
