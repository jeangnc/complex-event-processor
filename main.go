package main

import (
	"jeangnc/event-stream-filter/pkg/persistency/memory"
	"jeangnc/event-stream-filter/pkg/server"
)

const (
	port = ":8080"
)

func main() {
	a := memory.NewMemoryAdapter()
	s := server.NewGrpcServer(a)
	s.Start(port)
}
