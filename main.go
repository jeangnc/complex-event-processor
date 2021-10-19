package main

import (
	"encoding/json"
	"fmt"
	"jeangnc/pattern-matcher/pkg/tree"
	"sort"
)

type Event struct {
	Kind    string
	Payload map[string]string
}

type Predicate struct {
	Id              string            `json:"id"`
	ExpectedPayload map[string]string `json:"expected_payload"`
}

func mapKeys(hashmap map[string]string) []string {
	keys := make([]string, 0, len(hashmap))

	for k := range hashmap {
		keys = append(keys, k)
	}

	return keys
}

func main() {
	event := Event{
		Kind: "visit",
		Payload: map[string]string{
			"campo inutil": "qualquer coisa",
			"taitle":       "jndnka",
			"title":        "jasdjnkad",
			"url":          "/contato",
		},
	}

	p1 := &Predicate{
		Id: "1",
		ExpectedPayload: map[string]string{
			"taitle": "contato",
			"title":  "contato",
		},
	}

	p2 := &Predicate{
		Id: "2",
		ExpectedPayload: map[string]string{
			"taitle": "contato",
			"title":  "contato",
			"url":    "/contato",
		},
	}

	p3 := &Predicate{
		Id: "3",
		ExpectedPayload: map[string]string{
			"url": "/produtos",
		},
	}

	p4 := &Predicate{
		Id: "4",
		ExpectedPayload: map[string]string{
			"title": "contato",
		},
	}

	p5 := &Predicate{
		Id: "5",
		ExpectedPayload: map[string]string{
			"title": "contato",
			"url":   "/contato",
		},
	}

	tree := tree.NewTree()

	predicates := []*Predicate{p1, p2, p3, p4, p5}
	for _, predicate := range predicates {
		keys := mapKeys(predicate.ExpectedPayload)
		sort.Strings(keys)

		tree.Append(keys, predicate)
	}

	s, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println("tree", string(s))

	keys := mapKeys(event.Payload)
	sort.Strings(keys)

	foundPredicates := tree.Search(keys)
	for _, predicate := range foundPredicates {
		fmt.Println("search result", *(predicate.(*Predicate)))
	}
}
