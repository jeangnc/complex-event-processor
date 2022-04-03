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

func (n *Node) Set(key string, value interface{}) {
	n.value[key] = value
}

func (n *Node) Get(key string) interface{} {
	return n.value[key]
}

func (n *Node) Unset(key string) {
	delete(n.value, key)
}

func (n *Node) Traverse(path []string) *Node {
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

// traverses different arrangements of keys retrieving values along the way
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
