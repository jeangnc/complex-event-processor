package processor

import (
	"reflect"
	"testing"
)

// Tests entity and impact capability to merge
func TestImpactMerge(t *testing.T) {
	e := Entity{
		predicates: map[string]bool{
			"test1": true,
			"test2": true,
		},
	}

	i := Impact{
		predicates: map[string]bool{
			"test2": false,
			"test3": true,
		},
	}

	e1 := Process(e, i)
	expected := map[string]bool{
		"test1": true,
		"test2": false,
		"test3": true,
	}

	if !reflect.DeepEqual(e1.predicates, expected) {
		t.Fatalf(`predicates list is different than expected: %v`, e1.predicates)
	}

}

// Ensure an entity cannot mutate
func TestEntityImmutabilitty(t *testing.T) {
	e := Entity{
		predicates: map[string]bool{},
	}

	i := Impact{
		predicates: map[string]bool{
			"anything": true,
		},
	}

	Process(e, i)

	e1 := Entity{
		predicates: map[string]bool{},
	}

	if !reflect.DeepEqual(e, e1) {
		t.Fatalf(`original entity changed: %v`, e)
	}
}
