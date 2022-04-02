package tree

type Node struct {
	value ValueMap
	nodes NodeMap
}
type ValueMap map[string]interface{}
type NodeMap map[string]Node

func NewNode() Node {
	return Node{
		value: ValueMap{},
		nodes: NodeMap{},
	}
}

func (n *Node) Append(path []string, key string, value interface{}) {
	n.find(path).value[key] = value
}

func (n *Node) Remove(path []string, key string) {
	delete(n.find(path).value, key)
}

func (n Node) Values(path []string) []interface{} {
	var found []interface{}

	if n.value != nil {
		for _, v := range n.value {
			found = append(found, v)
		}
	}

	for i, k := range path {
		c, ok := n.nodes[k]

		if ok {
			found = append(found, c.Values(path[i:])...)
		}
	}

	return found
}

func (n *Node) find(path []string) *Node {
	target := n

	for _, k := range path {
		next, ok := target.nodes[k]

		if !ok {
			newNode := NewNode()
			target.nodes[k] = newNode
			next = newNode
		}

		target = &next
	}

	return target
}
