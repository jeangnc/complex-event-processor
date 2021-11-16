package main

import (
	"jeangnc/event-stream-filter/pkg/persistency"
	"jeangnc/event-stream-filter/pkg/persistency/filesystem"
	"jeangnc/event-stream-filter/pkg/server"
	"log"
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

	s := server.NewGrpcServer()
	s.Start(port)
}
