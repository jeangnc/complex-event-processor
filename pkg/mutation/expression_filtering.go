package mutation

type Changes struct {
	predicates map[string]bool
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
