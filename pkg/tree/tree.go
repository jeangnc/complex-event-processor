package tree

import (
	"sort"
)

type Node struct {
	children map[string]*Node
	values   []interface{}
}

func NewTree() *Node {
	return createEmptyNode()
}

func (node *Node) Append(keys []string, item interface{}) {
	sort.Strings(keys)

	var childNode, currentNode *Node
	var ok bool

	currentNode = node

	for _, key := range keys {
		if childNode, ok = currentNode.children[key]; !ok {
			childNode = createEmptyNode()
			currentNode.children[key] = childNode
		}

		currentNode = childNode
	}

	currentNode.values = append(currentNode.values, item)
}

func (node *Node) Search(keys []string) []interface{} {
	sort.Strings(keys)

	var foundPredicates []interface{}
	foundPredicates = append(foundPredicates, node.values...)

	for i, key := range keys {
		if child, ok := node.children[key]; ok {
			result := child.Search(keys[i+1:])
			foundPredicates = append(foundPredicates, result...)
		}
	}

	return foundPredicates
}

func createEmptyNode() *Node {
	return &Node{
		children: make(map[string]*Node),
		values:   make([]interface{}, 0, 0),
	}
}
