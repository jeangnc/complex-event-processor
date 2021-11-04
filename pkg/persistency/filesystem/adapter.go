package filesystem

import (
	"encoding/json"
	"jeangnc/event-stream-filter/pkg/tree"
	"jeangnc/event-stream-filter/pkg/types"
	"log"
	"os"
	"runtime"
)

type FilesystemAdapter struct {
	Filename string
}

func NewFilesystemAdapter(filename string) *FilesystemAdapter {
	return &FilesystemAdapter{
		Filename: filename,
	}
}

func (a *FilesystemAdapter) Load() *tree.ConditionTree {
	f, err := os.Open(a.Filename)

	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", a.Filename, err.Error())
	}

	t := tree.NewConditionTree()
	d := json.NewDecoder(f)

	for d.More() {
		c := types.Condition{}
		d.Decode(&c)
		t.Append(c)
	}

	// FIXME: remove after tuning memory footprint
	runtime.GC()

	return t
}
