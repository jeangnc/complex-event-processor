package persistency

import "jeangnc/event-stream-filter/pkg/tree"

type Adapter interface {
	Load() *tree.ConditionTree
}
