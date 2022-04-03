package tree

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	root := NewNode()
	node := root.Traverse([]string{"k1", "k2"})

	value := "my-string"
	node.Set("1", value)

	nodeValue := node.Get("1")

	if !reflect.DeepEqual(nodeValue, value) {
		t.Fatalf(`Failed to append: %v %v`, nodeValue, value)
	}
}

func TestUnset(t *testing.T) {
	root := Node{
		nodes: NodeMap{
			"k1": Node{
				value: ValueMap{"1": "my-string"},
			},
		},
	}

	node := root.Traverse([]string{"k1"})
	node.Unset("1")

	value := node.Get("1")

	if value != nil {
		t.Fatalf(`Failed to remove: %v`, value)
	}
}

// ensures correct values fetch over different keys arrangements
func TestValues(t *testing.T) {
	n := Node{
		nodes: NodeMap{
			"k1": Node{
				value: ValueMap{"1": "a"},
				nodes: NodeMap{
					"k2": Node{
						value: ValueMap{"2": "b"},
					},
				},
			},
			"k3": Node{
				value: ValueMap{"3": "c"},
			},
		},
	}

	found := n.Values([]string{"k1", "k2", "k3"})
	expected := []interface{}{"a", "b", "c"}

	if !reflect.DeepEqual(found, expected) {
		t.Fatalf(`Failed to retrieve all values: %v %v`, found, expected)
	}
}
