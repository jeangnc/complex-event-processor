package persistency

import (
	"encoding/json"
	"fmt"
	"jeangnc/pattern-matcher/pkg/tree"
	"jeangnc/pattern-matcher/pkg/types"
	"log"
	"os"
	"runtime"
)

func LoadFile(filename string) *tree.ConditionTree {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", filename, err.Error())
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
	printMemUsage()

	return t
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Printf("\n")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
