package main

import (
	"context"

	pb "github.com/jeangnc/complex-event-processor/pkg/proto"
	"github.com/jeangnc/complex-event-processor/pkg/server"
	"github.com/jeangnc/complex-event-processor/pkg/tree"
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

	// TODO: keep and index of expressions
	// TODO: persist expressions

	return &pb.RegisterResponse{}, nil
}

func (s *ServerImpl) Process(ctx context.Context, in *pb.ProcessRequest) (*pb.ProcessResponse, error) {
	conditions := s.tree.Search(in.Event)

	// TODO: impact processing
	// TODO: persist entities
	// TODO: notifications as response?

	return &pb.ProcessResponse{
		Event:      in.Event,
		Conditions: conditions,
	}, nil
}

// TODO: async API
// TODO: sql interface

func processEvent() {
	// load entity
	// iterates over impact evaluation result
	// mutates the entity
	// compares the previous and next result
	// emit notification
}

func main() {
	t := tree.NewConditionTree()
	i := &ServerImpl{
		tree: t,
	}

	s := server.NewGrpcServer(i)
	s.Start(port)
}
