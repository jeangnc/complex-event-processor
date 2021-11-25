package persistency

import (
	pb "jeangnc/event-stream-filter/pkg/proto"
)

type Adapter interface {
	Append(c *pb.Condition)
	Search(e *pb.Event) []*pb.Condition
}
