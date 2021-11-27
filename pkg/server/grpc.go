package server

import (
	"log"
	"net"

	pb "github.com/jeangnc/event-stream-filter/pkg/proto"

	"google.golang.org/grpc"
)

type grpcServer struct {
	serverImpl pb.EventStreamServer
}

func NewGrpcServer(s pb.EventStreamServer) *grpcServer {
	return &grpcServer{
		serverImpl: s,
	}
}

func (s *grpcServer) Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g := grpc.NewServer()
	pb.RegisterEventStreamServer(g, s.serverImpl)

	log.Printf("server listening at %v", lis.Addr())

	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
