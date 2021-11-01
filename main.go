package main

import (
	"flag"
	"fmt"
	"jeangnc/pattern-matcher/pkg/serialization"
	"log"
	"time"
)

func main() {
	filename := flag.String("filename", "", "The file containing a list of conditions")
	flag.Parse()

	if *filename == "" {
		log.Fatalf("A filename must be provided")
	}

	start := time.Now()
	fmt.Println("Loading tree")
	serialization.LoadJsonFile(*filename)
	fmt.Println("Initialization time:", time.Since(start))

	time.Sleep(60 * time.Second)
}
