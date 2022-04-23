package expression

const CONNECTOR_AND string = "and"
const CONNECTOR_OR string = "or"

func FilterImpacted(c Changes, es []Expression) []Expression {
	r := make([]Expression, 0, 0)

	for _, e := range es {
		for _, p := range e.predicates {
			if _, ok := c.predicates[p.Id]; ok {
				r = append(r, e)
			}
		}
	}

	return r
}
