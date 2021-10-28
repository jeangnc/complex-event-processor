package types

type Event struct {
	TenantId string
	Kind     string
	Payload  map[string]string
}

type Condition struct {
	Id            string `json:"id"`
	TenantId      string
	EventType     string
	Predicates    []Predicate `json:"predicates"`
	DesiredResult *bool       `json:"desired_result"`
}

type Predicate struct {
	Name     string      `json:"name"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}
