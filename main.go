package main

import (
	"context"
	"jeangnc/event-stream-filter/pkg/persistency"
	pb "jeangnc/event-stream-filter/pkg/proto"
	"jeangnc/event-stream-filter/pkg/server"
)

const (
	port = ":8080"
)

type ServerImpl struct {
	pb.UnimplementedEventStreamServer
	adapter persistency.Adapter
}

func (s *ServerImpl) RegisterCondition(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.adapter.Append(in.Condition)

	return &pb.RegisterResponse{}, nil
}

func (s *ServerImpl) Filter(ctx context.Context, in *pb.FilterRequest) (*pb.FilterResponse, error) {
	conditions := s.adapter.Search(in.Event)

	return &pb.FilterResponse{
		Event:      in.Event,
		Conditions: conditions,
	}, nil
}

func main() {
	a := persistency.NewMemoryAdapter()
	i := &ServerImpl{
		adapter: a,
	}

	s := server.NewGrpcServer(i)
	s.Start(port)
}
