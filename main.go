package main

import (
	"fmt"
	"jeangnc/event-stream-filter/pkg/grpc"
	"jeangnc/event-stream-filter/pkg/persistency"
	"jeangnc/event-stream-filter/pkg/persistency/filesystem"
	"log"
	"runtime"
)

const (
	port = ":8080"
)

type Options struct {
	filename string
}

func buildAdapter(o Options) persistency.Adapter {
	if o.filename == "" {
		log.Fatalf("A filename must be provided")
	}

	return filesystem.NewFilesystemAdapter(o.filename)
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

func main() {
	/*
		filename := flag.String("filename", "", "The file containing a list of conditions")
		flag.Parse()

		options := Options{filename: *filename}
		adapter := buildAdapter(options)

		start := time.Now()
		fmt.Println("Loading tree")
		adapter.Load()
		fmt.Println("Initialization time:", time.Since(start))
	*/

	s := grpc.NewServer()
	s.Start(port)
}
