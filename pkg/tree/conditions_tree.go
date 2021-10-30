package tree

import (
	"jeangnc/pattern-matcher/pkg/types"
)

type ConditionTree struct {
	treeIndex map[string]map[string]*Node
}

func NewConditionTree() *ConditionTree {
	return &ConditionTree{
		treeIndex: make(map[string]map[string]*Node),
	}
}

func (conditionTree *ConditionTree) Append(condition types.Condition) {
	keys := make([]string, 0, len(condition.Predicates))

	for _, predicate := range condition.Predicates {
		keys = append(keys, predicate.Name)
	}

	eventTypeIndex, ok := conditionTree.treeIndex[condition.TenantId]
	if !ok {
		eventTypeIndex = make(map[string]*Node)
		conditionTree.treeIndex[condition.TenantId] = eventTypeIndex
	}

	eventTree, ok := eventTypeIndex[condition.EventType]
	if !ok {
		eventTree = NewTree()
		eventTypeIndex[condition.EventType] = eventTree
	}

	eventTree.Append(keys, condition)
}

func (conditionTree *ConditionTree) AppendMultiple(conditions []types.Condition) {
	for _, condition := range conditions {
		conditionTree.Append(condition)
	}
}

func (conditionTree *ConditionTree) Search(event types.Event) []*types.Condition {
	foundConditions := make([]*types.Condition, 0, 0)
	tree := conditionTree.findTree(event.TenantId, event.Kind)

	if tree == nil {
		return foundConditions
	}

	payloadKeys := extractKeys(event.Payload)
	foundNodes := tree.Search(payloadKeys)

	for _, node := range foundNodes {
		condition := node.(types.Condition)
		evaluationResult := evaluateCondition(condition, event)

		if condition.DesiredResult == nil || evaluationResult == *condition.DesiredResult {
			foundConditions = append(foundConditions, &condition)
		}
	}

	return foundConditions
}

func (conditionTree *ConditionTree) findTree(tenantId string, eventType string) *Node {
	eventTypeIndex, ok := conditionTree.treeIndex[tenantId]
	if !ok {
		return nil
	}

	eventTree, ok := eventTypeIndex[eventType]
	if !ok {
		return nil
	}

	return eventTree
}

func extractKeys(hashmap map[string]interface{}) []string {
	keys := make([]string, 0, len(hashmap))

	for k := range hashmap {
		keys = append(keys, k)
	}

	return keys
}

func evaluateCondition(condition types.Condition, event types.Event) bool {
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
