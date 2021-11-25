package server

import (
	"context"
	pb "jeangnc/event-stream-filter/pkg/proto"
	"jeangnc/event-stream-filter/pkg/tree"
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedEventStreamServer

	tree *tree.ConditionTree
}

func NewGrpcServer() *grpcServer {
	t := tree.NewConditionTree()

	return &grpcServer{
		tree: t,
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
	s.tree.Append(in.Condition)

	return &pb.RegisterResponse{}, nil
}

func (s *grpcServer) Filter(ctx context.Context, in *pb.FilterRequest) (*pb.FilterResponse, error) {
	conditions := s.tree.Search(in.Event)

	return &pb.FilterResponse{
		Event:      in.Event,
		Conditions: conditions,
	}, nil
}
