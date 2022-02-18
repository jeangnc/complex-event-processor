package processing

import (
	util "github.com/jeangnc/complex-event-processor/pkg/util"
)

func Process(e Entity, i Impact) Entity {
	return Entity{
		predicates: util.MergeMaps(e.predicates, i.predicates),
	}
}
