package types

type LogicalExpression struct {
	Connector string    `json:"connector"`
	Operands  []Operand `json:"operands"`
}

func (le LogicalExpression) Predicates() []*Predicate {
	r := make([]*Predicate, 0)

	for _, o := range le.Operands {
		if o.Predicate != nil {
			r = append(r, o.Predicate)
			continue
		}

		r = append(r, o.LogicalExpression.Predicates()...)
	}

	return r
}
