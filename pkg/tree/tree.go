package tree

import "sort"

type Node struct {
	Nodes map[string]*Node `json:"nodes"`
	Items []interface{}    `json:"items"`
}

func NewTree() *Node {
	return createEmptyNode()
}

func (tree *Node) Append(keys []string, item interface{}) {
	sort.Strings(keys)
	currentNode := tree

	for _, key := range keys {
		var node *Node
		var ok bool

		if node, ok = currentNode.Nodes[key]; !ok {
			node = createEmptyNode()
			currentNode.Nodes[key] = node
		}

		currentNode = node
	}

	currentNode.Items = append(currentNode.Items, item)
}

func (tree *Node) Search(keys []string) []interface{} {
	sort.Strings(keys)

	var foundPredicates []interface{}
	foundPredicates = append(foundPredicates, tree.Items...)

	for i, key := range keys {
		if subtree, ok := tree.Nodes[key]; ok {
			result := subtree.Search(keys[i+1:])
			foundPredicates = append(foundPredicates, result...)
		}
	}

	return foundPredicates
}

func createEmptyNode() *Node {
	return &Node{
		Nodes: make(map[string]*Node),
		Items: make([]interface{}, 0),
	}
}
