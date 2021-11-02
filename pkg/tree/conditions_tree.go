package tree

import (
	"jeangnc/pattern-matcher/pkg/types"
)

type ConditionTree struct {
	tenantIndex map[string]map[string]*Node
}

func NewConditionTree() *ConditionTree {
	return &ConditionTree{
		tenantIndex: make(map[string]map[string]*Node),
	}
}

func (c *ConditionTree) Append(condition types.Condition) {
	keys := make([]string, 0, len(condition.Predicates))

	for _, predicate := range condition.Predicates {
		keys = append(keys, predicate.Name)
	}

	eventTree := c.findTree(condition.TenantId, condition.EventType, true)
	eventTree.Append(keys, &condition)
}

func (c *ConditionTree) AppendMultiple(conditions []types.Condition) {
	for _, condition := range conditions {
		c.Append(condition)
	}
}

func (c *ConditionTree) Search(event types.Event) []*types.Condition {
	foundConditions := make([]*types.Condition, 0, 0)

	tree := c.findTree(event.TenantId, event.Kind, false)
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

func (c *ConditionTree) findTree(tenantId string, eventType string, fill bool) *Node {
	eventTypeIndex, ok := c.tenantIndex[tenantId]
	if !ok {
		if !fill {
			return nil
		}
		eventTypeIndex = make(map[string]*Node)
		c.tenantIndex[tenantId] = eventTypeIndex
	}

	eventTree, ok := eventTypeIndex[eventType]
	if !ok {
		if !fill {
			return nil
		}
		eventTree = NewTree()
		eventTypeIndex[eventType] = eventTree
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
