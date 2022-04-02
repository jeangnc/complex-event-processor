package tree

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	node := NewNode()
	value := "my-string"

	node.Append([]string{"k1", "k2"}, "1", value)

	nodeValue := node.find([]string{"k1", "k2"}).value["1"]

	if !reflect.DeepEqual(nodeValue, value) {
		t.Fatalf(`Failed to append: %v %v`, nodeValue, value)
	}
}

func TestRemove(t *testing.T) {
	node := Node{
		nodes: NodeMap{
			"k1": Node{
				value: ValueMap{"1": "my-string"},
			},
		},
	}

	node.Remove([]string{"k1"}, "1")

	value := node.nodes["k1"].value["1"]

	if value != nil {
		t.Fatalf(`Failed to remove: %v`, value)
	}
}

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
