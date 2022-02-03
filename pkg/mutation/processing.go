package mutation

type Entity struct {
	predicates map[string]bool
	relations  map[string]bool
}

type Impact struct {
	predicates map[string]bool
	relations  map[string]bool
}

func Process(e Entity, i Impact) Entity {
	return Entity{
		predicates: mergeMap(e.predicates, i.predicates),
		relations:  mergeMap(e.relations, i.relations),
	}
}

func mergeMap(m1 map[string]bool, m2 map[string]bool) map[string]bool {
	newMap := copyMap(m1)
	for k, v := range m2 {
		newMap[k] = v
	}

	return newMap
}

func copyMap(m map[string]bool) map[string]bool {
	newMap := map[string]bool{}

	for k, v := range m {
		newMap[k] = v
	}

	return newMap
}
