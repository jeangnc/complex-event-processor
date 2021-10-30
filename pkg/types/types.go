package types

type Event struct {
	Id       string                 `json:"id"`
	TenantId string                 `json:"tenant_id"`
	Kind     string                 `json:"kind"`
	Payload  map[string]interface{} `json:"payload"`
}

type Condition struct {
	Id            string      `json:"id"`
	TenantId      string      `json:"tenant_id"`
	EventType     string      `json:"event_type"`
	Predicates    []Predicate `json:"predicates"`
	DesiredResult *bool       `json:"desired_result"`
}

type Predicate struct {
	Name     string      `json:"name"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}
