package mutation

import (
	"github.com/jeangnc/complex-event-processor/pkg/types"
	util "github.com/jeangnc/complex-event-processor/pkg/util"
)

func Process(e types.Entity, i types.Impact) (types.Entity, types.Changes) {
	mergedPredicates, changedKeys := util.MergeMaps(e.Predicates, i.Predicates)
	newEntity := types.Entity{Predicates: mergedPredicates}

	changed := map[string]bool{}
	for _, k := range changedKeys {
		changed[k] = true
	}
	changes := types.Changes{Predicates: changed}

	return newEntity, changes
}
