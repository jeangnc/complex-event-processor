package main

import (
	"context"

	pb "github.com/jeangnc/event-stream-filter/pkg/proto"
	"github.com/jeangnc/event-stream-filter/pkg/server"
	"github.com/jeangnc/event-stream-filter/pkg/tree"
)

const (
	port = ":8080"
)

type ServerImpl struct {
	pb.UnimplementedEventStreamServer
	tree *tree.ConditionTree
}

func (s *ServerImpl) RegisterCondition(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.tree.Append(in.Condition)

	return &pb.RegisterResponse{}, nil
}

func (s *ServerImpl) Filter(ctx context.Context, in *pb.FilterRequest) (*pb.FilterResponse, error) {
	conditions := s.tree.Search(in.Event)

	return &pb.FilterResponse{
		Event:      in.Event,
		Conditions: conditions,
	}, nil
}

func main() {
	t := tree.NewConditionTree()
	i := &ServerImpl{
		tree: t,
	}

	s := server.NewGrpcServer(i)
	s.Start(port)
}
