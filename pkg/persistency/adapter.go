package persistency

import "jeangnc/pattern-matcher/pkg/tree"

type Adapter interface {
	Load() *tree.ConditionTree
}
