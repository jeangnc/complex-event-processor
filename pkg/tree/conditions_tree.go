package tree

import (
	"fmt"
	"jeangnc/pattern-matcher/pkg/types"
)

type ConditionTree struct {
	Tree *Node `json:"tree"`
}

func NewConditionTree() *ConditionTree {
	return &ConditionTree{
		Tree: NewTree(),
	}
}

func extractKeys(hashmap map[string]string) []string {
	keys := make([]string, 0, len(hashmap))

	for k := range hashmap {
		keys = append(keys, k)
	}

	return keys
}

func evaluateCondition(condition *types.Condition, event *types.Event) bool {
	result := true

	for _, predicate := range condition.Predicates {
		payloadValue := event.Payload[predicate.Name]

		switch predicate.Operator {
		case "equal":
			result = result && predicate.Value == payloadValue
		default:
			result = false
		}
	}

	return result
}

func (conditionTree *ConditionTree) Append(conditions []*types.Condition) {
	for _, condition := range conditions {
		keys := make([]string, 0, len(condition.Predicates))

		for _, predicate := range condition.Predicates {
			keys = append(keys, predicate.Name)
		}

		conditionTree.Tree.Append(keys, condition)
	}
}

func (conditionTree *ConditionTree) Search(event *types.Event) []*types.Condition {
	payloadKeys := extractKeys(event.Payload)

	foundNodes := conditionTree.Tree.Search(payloadKeys)
	foundConditions := make([]*types.Condition, 0, 0)

	for _, node := range foundNodes {
		condition := node.(*types.Condition)
		evaluationResult := evaluateCondition(condition, event)

		fmt.Println(condition, evaluationResult)

		if evaluationResult {
			foundConditions = append(foundConditions, condition)
		}
	}

	return foundConditions
}
