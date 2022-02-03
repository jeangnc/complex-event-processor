package processor

import (
	"reflect"
	"testing"
)

// Tests entity and impact capability to merge
func TestImpactMerge(t *testing.T) {
	e := Entity{
		predicates: []string{
			"test1",
			"test2",
		},
	}

	i := Impact{
		predicates: map[string]bool{
			"test2": false,
			"test3": true,
		},
	}

	e1 := Process(e, i)
	expected := []string{
		"test1",
		"test3",
	}

	if !reflect.DeepEqual(e1.predicates, expected) {
		t.Fatalf(`predicates list is different than expected: %v`, e1.predicates)
	}

}

// Ensure an entity cannot mutate
func TestEntityImmutabilitty(t *testing.T) {
	e := Entity{}

	i := Impact{
		predicates: map[string]bool{
			"anything": true,
		},
	}

	Process(e, i)

	if !reflect.DeepEqual(e, Entity{}) {
		t.Fatalf(`original entity changed`)
	}
}
