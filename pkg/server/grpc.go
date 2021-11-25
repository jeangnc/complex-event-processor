package server

import (
	"context"
	"jeangnc/event-stream-filter/pkg/persistency"
	pb "jeangnc/event-stream-filter/pkg/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedEventStreamServer

	adapter persistency.Adapter
}

func NewGrpcServer(a persistency.Adapter) *grpcServer {
	return &grpcServer{
		adapter: a,
	}
}

func (s *grpcServer) Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g := grpc.NewServer()
	pb.RegisterEventStreamServer(g, s)

	log.Printf("server listening at %v", lis.Addr())

	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *grpcServer) RegisterCondition(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.adapter.Append(in.Condition)

	return &pb.RegisterResponse{}, nil
}

func (s *grpcServer) Filter(ctx context.Context, in *pb.FilterRequest) (*pb.FilterResponse, error) {
	conditions := s.adapter.Search(in.Event)

	return &pb.FilterResponse{
		Event:      in.Event,
		Conditions: conditions,
	}, nil
}
