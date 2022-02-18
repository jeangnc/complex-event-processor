package processing

type Expression struct {
	predicates        []string
	logicalExpression *LogicalExpression
}

func FilterImpacted(c Changes, es []Expression) []Expression {
	r := make([]Expression, 0, 0)

	for _, e := range es {
		for _, p := range e.predicates {
			if _, ok := c.predicates[p]; ok {
				r = append(r, e)
			}
		}
	}

	return r
}
