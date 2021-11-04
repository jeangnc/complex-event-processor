package main

import (
	"context"
	"fmt"
	"jeangnc/event-stream-filter/pkg/persistency"
	"jeangnc/event-stream-filter/pkg/persistency/filesystem"
	pb "jeangnc/event-stream-filter/pkg/proto"
	"log"
	"net"
	"runtime"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type Options struct {
	filename string
}

type server struct {
	pb.UnimplementedEventStreamServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
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

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEventStreamServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
