package types

import (
	"reflect"
	"testing"
)

func TestPredicates(t *testing.T) {
	a := Predicate{Id: "a"}
	b := Predicate{Id: "b"}

	le := LogicalExpression{
		Connector: CONNECTOR_AND,
		Operands: []Operand{
			Operand{Predicate: &a},
			Operand{Predicate: &b},
		},
	}

	predicates := le.DeepPredicates()
	expected := []*Predicate{&a, &b}

	if !reflect.DeepEqual(predicates, expected) {
		t.Fatalf(`Failed to fetch predicates list: %v %v`, predicates, expected)
	}
}
