package main

import (
	"fmt"
	"jeangnc/pattern-matcher/pkg/tree"
	"jeangnc/pattern-matcher/pkg/types"
)

func initializeTree() *tree.ConditionTree {
	var desiredResult bool

	desiredResult = true

	c1 := &types.Condition{
		Id:            "1",
		DesiredResult: &desiredResult,
		Predicates: []types.Predicate{
			types.Predicate{
				Name:     "origin",
				Operator: "equal",
				Value:    "2",
			},
			types.Predicate{
				Name:     "title",
				Operator: "equal",
				Value:    "contact",
			},
		},
	}

	c2 := &types.Condition{
		Id:            "2",
		DesiredResult: &desiredResult,
		Predicates: []types.Predicate{
			types.Predicate{
				Name:     "origin",
				Operator: "equal",
				Value:    "3",
			},
			types.Predicate{
				Name:     "title",
				Operator: "equal",
				Value:    "contact",
			},
			types.Predicate{
				Name:     "url",
				Operator: "equal",
				Value:    "/contact",
			},
		},
	}

	c3 := &types.Condition{
		Id:            "3",
		DesiredResult: &desiredResult,
		Predicates: []types.Predicate{
			types.Predicate{
				Name:     "url",
				Operator: "equal",
				Value:    "/products",
			},
		},
	}

	c4 := &types.Condition{
		Id:            "4",
		DesiredResult: &desiredResult,
		Predicates: []types.Predicate{
			types.Predicate{
				Name:     "title",
				Operator: "equal",
				Value:    "contact",
			},
		},
	}

	c5 := &types.Condition{
		Id: "5",
		Predicates: []types.Predicate{
			types.Predicate{
				Name:     "title",
				Operator: "equal",
				Value:    "contact",
			},
			types.Predicate{
				Name:     "url",
				Operator: "equal",
				Value:    "/contact",
			},
		},
	}

	tree := tree.NewConditionTree()
	conditions := []*types.Condition{c1, c2, c3, c4, c5}
	tree.Append(conditions)

	return tree
}

func main() {
	event := &types.Event{
		Kind: "visit",
		Payload: map[string]string{
			"useless": "anything",
			"url":     "/contact",
			"title":   "contact",
			"origin":  "2",
		},
	}

	tree := initializeTree()

	conditions := tree.Search(event)

	for _, condition := range conditions {
		fmt.Println(condition)
	}
}
