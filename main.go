package main

import (
	"flag"
	"fmt"
	"jeangnc/pattern-matcher/pkg/persistency/filesystem"
	"log"
	"runtime"
	"time"
)

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
	filename := flag.String("filename", "", "The file containing a list of conditions")
	flag.Parse()

	if *filename == "" {
		log.Fatalf("A filename must be provided")
	}

	start := time.Now()
	fmt.Println("Loading tree")
	filesystem.Load(*filename)
	fmt.Println("Initialization time:", time.Since(start))

	printMemUsage()

	time.Sleep(60 * time.Second)
}
