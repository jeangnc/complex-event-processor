package filesystem

import (
	"encoding/json"
	pb "jeangnc/event-stream-filter/pkg/proto"
	"jeangnc/event-stream-filter/pkg/tree"
	"log"
	"os"
	"runtime"
)

type filesystemAdapter struct {
	tree     *tree.ConditionTree
	filename string
}

func NewFilesystemAdapter(filename string) *filesystemAdapter {
	return &filesystemAdapter{
		tree:     tree.NewConditionTree(),
		filename: filename,
	}
}

func (a *filesystemAdapter) Load() *tree.ConditionTree {
	f, err := os.Open(a.filename)

	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", a.filename, err.Error())
	}

	d := json.NewDecoder(f)

	for d.More() {
		c := pb.Condition{}
		d.Decode(&c)
		a.tree.Append(&c)
	}

	// FIXME: remove after tuning memory footprint
	runtime.GC()

	return t
}

func (a *filesystemAdapter) Append(c *pb.Condition) {
	// TODO: implement
}
