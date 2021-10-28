package types

type Event struct {
	Kind    string
	Payload map[string]string
}

type Condition struct {
	Id            string      `json:"id"`
	Predicates    []Predicate `json:"predicates"`
	DesiredResult bool
}

type Predicate struct {
	Name     string      `json:"name"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}
