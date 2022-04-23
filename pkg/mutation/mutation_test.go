package mutation

import (
	"reflect"
	"testing"

	"github.com/jeangnc/complex-event-processor/pkg/types"
)

// Tests entity and impact capability to merge
func TestImpactMerge(t *testing.T) {
	e := types.Entity{
		Predicates: map[string]bool{
			"test1": true,
			"test2": true,
		},
	}

	i := types.Impact{
		Predicates: map[string]bool{
			"test2": false,
			"test3": true,
		},
	}

	e1, changes := Process(e, i)
	expected := map[string]bool{
		"test1": true,
		"test2": false,
		"test3": true,
	}

	if !reflect.DeepEqual(e1.Predicates, expected) {
		t.Fatalf(`predicates list is different than expected: %v`, e1.Predicates)
	}

	expectedChanges := types.Changes{
		Predicates: map[string]bool{
			"test2": true,
			"test3": true,
		},
	}
	if !reflect.DeepEqual(changes, expectedChanges) {
		t.Fatalf(`changes list is different than expected: %v %v`, changes, expectedChanges)
	}
}

// Ensure entity doesnt mutate
func TestEntityImmutabilitty(t *testing.T) {
	e := types.Entity{
		Predicates: map[string]bool{},
	}

	i := types.Impact{
		Predicates: map[string]bool{
			"anything": true,
		},
	}

	Process(e, i)

	e1 := types.Entity{
		Predicates: map[string]bool{},
	}

	if !reflect.DeepEqual(e, e1) {
		t.Fatalf(`original entity changed: %v`, e)
	}
}
