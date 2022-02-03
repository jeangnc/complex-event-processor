package mutation

// TODO: needing to specify the predicate entity type
// TODO: need to specify the related entity type

type Changes struct {
	predicates map[string]bool
	relations  map[string]bool
}

type Expression struct {
	predicates []string
}

func Impacted(c Changes, e Expression) bool {
	for _, p := range e.predicates {
		if _, ok := c.predicates[p]; ok {
			return true
		}
	}

	return false
}

// TODO: related entities share their state.
//  expressions whose predicate type matches the new related entity type must be impacted
