package grpc

import (
	"context"
	pb "jeangnc/event-stream-filter/pkg/proto"
	"log"
	"net"

	google_grpc "google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEventStreamServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Filter(ctx context.Context, in *pb.FilterRequest) (*pb.FilterResponse, error) {
	return &pb.FilterResponse{}, nil
}

func (s *server) Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := google_grpc.NewServer()
	pb.RegisterEventStreamServer(grpcServer, s)

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
