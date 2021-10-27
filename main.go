package main

import (
	"encoding/json"
	"fmt"
	tree "jeangnc/pattern-matcher/pkg/tree"
)

type Event struct {
	Kind    string
	Payload map[string]string
}

type Condition struct {
	Id         string      `json:"id"`
	Predicates []Predicate `json:"predicates"`
}

type Predicate struct {
	Name     string      `json:"name"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

func extractKeys(hashmap map[string]string) []string {
	keys := make([]string, 0, len(hashmap))

	for k := range hashmap {
		keys = append(keys, k)
	}

	return keys
}

func initializeTree() *tree.Node {
	c1 := &Condition{
		Id: "1",
		Predicates: []Predicate{
			Predicate{
				Name:     "origin",
				Operator: "equal",
				Value:    "2",
			},
			Predicate{
				Name:     "title",
				Operator: "equal",
				Value:    "contato",
			},
		},
	}

	c2 := &Condition{
		Id: "2",
		Predicates: []Predicate{
			Predicate{
				Name:     "origin",
				Operator: "equal",
				Value:    "3",
			},
			Predicate{
				Name:     "title",
				Operator: "equal",
				Value:    "contato",
			},
			Predicate{
				Name:     "url",
				Operator: "equal",
				Value:    "/contato",
			},
		},
	}

	c3 := &Condition{
		Id: "3",
		Predicates: []Predicate{
			Predicate{
				Name:     "url",
				Operator: "equal",
				Value:    "/produtos",
			},
		},
	}

	c4 := &Condition{
		Id: "4",
		Predicates: []Predicate{
			Predicate{
				Name:     "title",
				Operator: "equal",
				Value:    "contato",
			},
		},
	}

	c5 := &Condition{
		Id: "5",
		Predicates: []Predicate{
			Predicate{
				Name:     "title",
				Operator: "equal",
				Value:    "contato",
			},
			Predicate{
				Name:     "url",
				Operator: "equal",
				Value:    "/contato",
			},
		},
	}

	tree := tree.NewTree()

	conditions := []*Condition{c1, c2, c3, c4, c5}
	for _, condition := range conditions {
		keys := make([]string, 0, len(condition.Predicates))

		for _, predicate := range condition.Predicates {
			keys = append(keys, predicate.Name)
		}

		tree.Append(keys, condition)
	}

	s, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println("tree", string(s))

	return tree
}

func main() {
	event := Event{
		Kind: "visit",
		Payload: map[string]string{
			"useless": "anything",
			"url":     "/contact",
			"title":   "contact",
			"origin":  "2",
		},
	}

	tree := initializeTree()
	payloadKeys := extractKeys(event.Payload)

	foundNodes := tree.Search(payloadKeys)
	for _, node := range foundNodes {
		condition := *(node.(*Condition))

		fmt.Println("search result", condition)
	}
}
