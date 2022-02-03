package processor

type Entity struct {
	predicates []string
	relations  []string
}

type Impact struct {
	predicates map[string]bool
	relations  map[string]bool
}

func Process(e Entity, i Impact) Entity {

	return e
}
