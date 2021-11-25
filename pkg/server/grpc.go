package server

import (
	"context"
	"fmt"
	pb "jeangnc/event-stream-filter/pkg/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedEventStreamServer
}

func NewGrpcServer() *grpcServer {
	return &grpcServer{}
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
	fmt.Println("register request", in.Condition)
	return &pb.RegisterResponse{}, nil
}

func (s *grpcServer) Filter(ctx context.Context, in *pb.FilterRequest) (*pb.FilterResponse, error) {
	log.Printf("filter request %v", in)
	return &pb.FilterResponse{}, nil
}
