package main

import (
	"encoding/json"
	"fmt"
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

type Node struct {
	Nodes      map[string]Node `json:"nodes"`
	Predicates []Predicate     `json:"predicates"`
}

func buildTree(predicates []Predicate) Node {
	tree := createEmptyNode()

	var node Node
	var ok bool

	for _, predicate := range predicates {
		keys := mapKeys(predicate.ExpectedPayload)
		sort.Strings(keys)

		nodeIterator := tree
		for _, key := range keys {
			if node, ok = nodeIterator.Nodes[key]; !ok {
				node = createEmptyNode()
				nodeIterator.Nodes[key] = node
			}

			nodeIterator = node
		}
	}

	return tree
}

func createEmptyNode() Node {
	return Node{
		Nodes: make(map[string]Node),
	}
}

func searchRelevantPredicates(node Node, keys []string) []Predicate {
	var foundPredicates []Predicate

	foundPredicates = append(foundPredicates, node.Predicates...)

	for i, key := range keys {
		if subnode, ok := node.Nodes[key]; ok {
			result := searchRelevantPredicates(subnode, keys[i+1:])
			foundPredicates = append(foundPredicates, result...)
		}
	}

	return foundPredicates
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

	p1 := Predicate{
		Id: "1",
		ExpectedPayload: map[string]string{
			"taitle": "contato",
			"title":  "contato",
		},
	}

	p2 := Predicate{
		Id: "2",
		ExpectedPayload: map[string]string{
			"taitle": "contato",
			"title":  "contato",
			"url":    "/contato",
		},
	}

	p3 := Predicate{
		Id: "3",
		ExpectedPayload: map[string]string{
			"url": "/produtos",
		},
	}

	p4 := Predicate{
		Id: "4",
		ExpectedPayload: map[string]string{
			"title": "contato",
		},
	}

	p5 := Predicate{
		Id: "5",
		ExpectedPayload: map[string]string{
			"title": "contato",
			"url":   "/contato",
		},
	}

	tree := buildTree([]Predicate{p1, p2, p3, p4, p5})
	treeString, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(treeString))

	keys := mapKeys(event.Payload)
	sort.Strings(keys)

	predicates := searchRelevantPredicates(tree, keys)
	fmt.Println("ignore", predicates)
}
