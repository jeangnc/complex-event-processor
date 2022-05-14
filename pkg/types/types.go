package types

type Event struct {
	Id        string                 `json:"id"`
	TenantId  string                 `json:"tenant_id"`
	EntityId  string                 `json:"entity_id"`
	Type      string                 `json:"type"`
	Timestamp int64                  `json:"timestamp"`
	Payload   map[string]interface{} `json:"payload"`
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
