package tree

import (
	pb "jeangnc/event-stream-filter/pkg/proto"
)

type ConditionTree struct {
	tenantIndex map[string]map[string]*Node
}

func NewConditionTree() *ConditionTree {
	return &ConditionTree{
		tenantIndex: make(map[string]map[string]*Node),
	}
}

func (c *ConditionTree) Append(condition pb.Condition) {
	keys := make([]string, 0, len(condition.Predicates))

	for _, predicate := range condition.Predicates {
		keys = append(keys, predicate.Name)
	}

	eventTree := c.findTree(condition.TenantId, condition.EventType, true)
	eventTree.Append(keys, &condition)
}

func (c *ConditionTree) AppendMultiple(conditions []pb.Condition) {
	for _, condition := range conditions {
		c.Append(condition)
	}
}

func (c *ConditionTree) Search(event pb.Event) []*pb.Condition {
	foundConditions := make([]*pb.Condition, 0, 0)

	tree := c.findTree(event.TenantId, event.Kind, false)
	if tree == nil {
		return foundConditions
	}

	payload := event.Payload.AsMap()
	payloadKeys := extractKeys(payload)
	foundNodes := tree.Search(payloadKeys)

	for _, node := range foundNodes {
		condition := node.(pb.Condition)
		evaluationResult := evaluateCondition(condition, payload)

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

func evaluateCondition(condition pb.Condition, payload map[string]interface{}) bool {
	result := true

	for _, predicate := range condition.Predicates {
		payloadValue := payload[predicate.Name]

		switch predicate.Operator {
		case "equal":
			result = result && predicate.Value == payloadValue
		default:
			result = false
		}
	}

	return result
}
