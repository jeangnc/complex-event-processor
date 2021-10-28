package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"jeangnc/pattern-matcher/pkg/tree"
	"jeangnc/pattern-matcher/pkg/types"
	"log"
	"os"
	"time"
)

func initializeStubedTree() *tree.ConditionTree {
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
	tree.AppendMultiple(conditions)

	return tree
}

func importPredicatesFromFile(filename string) *tree.ConditionTree {
	tree := tree.NewConditionTree()

	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", filename, err.Error())
	}

	s := bufio.NewScanner(f)
	var jsonmap map[string]interface{}
	var data []byte

	for s.Scan() {
		data = []byte(s.Text())
		json.Unmarshal(data, &jsonmap)

		predicates := make([]types.Predicate, 0, 0)
		source := jsonmap["_source"].(map[string]interface{})

		textArguments := source["text_arguments"].([]interface{})
		for _, rawArgument := range textArguments {
			textArgument := rawArgument.(map[string]interface{})

			predicate := types.Predicate{
				Name:     textArgument["name"].(string),
				Operator: textArgument["operator"].(string),
				Value:    textArgument["value"].(string),
			}

			predicates = append(predicates, predicate)
		}

		numericArguments := source["numeric_arguments"].([]interface{})
		for _, rawArgument := range numericArguments {
			numericArgument := rawArgument.(map[string]interface{})

			predicate := types.Predicate{
				Name:     numericArgument["name"].(string),
				Operator: numericArgument["operator"].(string),
				Value:    numericArgument["value"].(float64),
			}

			predicates = append(predicates, predicate)
		}

		desiredResult := true
		condition := &types.Condition{
			Id:            jsonmap["_id"].(string),
			TenantId:      source["tenant_id"].(string),
			EventType:     source["object_type"].(string),
			DesiredResult: &desiredResult,
			Predicates:    predicates,
		}

		tree.Append(condition)
	}

	return tree
}

func main() {
	event := &types.Event{
		TenantId: "401000000000570740",
		Kind:     "CONVERSION",
		Payload: map[string]string{
			"conversion_identifier": "sorteio-geracao-de-valor",
		},
	}

	start := time.Now()
	tree := importPredicatesFromFile("./predicates.json")
	fmt.Println("Initialization time:", time.Since(start))

	// tree = initializeStubedTree()
	// bytes, _ := json.MarshalIndent(*tree.Tree, "", "  ")
	// fmt.Println(string(bytes))

	start = time.Now()
	conditions := tree.Search(event)
	fmt.Println("Search time:", time.Since(start))

	for _, condition := range conditions {
		fmt.Println(*condition)
	}

	// time.Sleep(60 * time.Second)
}
