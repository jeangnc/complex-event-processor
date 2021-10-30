package tree

import "sort"

type Node struct {
	nodes map[string]*Node
	items []interface{}
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

		if node, ok = currentNode.nodes[key]; !ok {
			node = createEmptyNode()
			currentNode.nodes[key] = node
		}

		currentNode = node
	}

	currentNode.items = append(currentNode.items, item)
}

func (tree *Node) Search(keys []string) []interface{} {
	sort.Strings(keys)

	var foundPredicates []interface{}
	foundPredicates = append(foundPredicates, tree.items...)

	for i, key := range keys {
		if subtree, ok := tree.nodes[key]; ok {
			result := subtree.Search(keys[i+1:])
			foundPredicates = append(foundPredicates, result...)
		}
	}

	return foundPredicates
}

func createEmptyNode() *Node {
	return &Node{
		nodes: make(map[string]*Node),
		items: make([]interface{}, 0, 0),
	}
}
