package mutation

import (
	util "github.com/jeangnc/complex-event-processor/pkg/util"
)

type Entity struct {
	predicates map[string]bool
}

type Impact struct {
	predicates map[string]bool
}

func Process(e Entity, i Impact) Entity {
	return Entity{
		predicates: util.MergeMap(e.predicates, i.predicates),
	}
}
