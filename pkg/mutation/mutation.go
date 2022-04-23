package processing

import (
	util "github.com/jeangnc/complex-event-processor/pkg/util"
)

func Process(e Entity, i Impact) (Entity, Changes) {
	mergedPredicates, changedKeys := util.MergeMaps(e.predicates, i.predicates)
	newEntity := Entity{predicates: mergedPredicates}

	changed := map[string]bool{}
	for _, k := range changedKeys {
		changed[k] = true
	}
	changes := Changes{predicates: changed}

	return newEntity, changes
}
