package memory

import (
	pb "jeangnc/event-stream-filter/pkg/proto"
	"jeangnc/event-stream-filter/pkg/tree"
)

// TODO: implement

type memoryAdapter struct {
	tree *tree.ConditionTree
}

func NewMemoryAdapter() *memoryAdapter {
	return &memoryAdapter{
		tree: tree.NewConditionTree(),
	}
}

func (a *memoryAdapter) Append(c *pb.Condition) {
	a.tree.Append(c)
}

func (a *memoryAdapter) Search(e *pb.Event) []*pb.Condition {
	return a.tree.Search(e)
}
