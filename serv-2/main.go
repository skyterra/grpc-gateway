package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "xframework/protobuf"
)

const (
	port = ":50052"
)

type server struct {
	pb.UnimplementedMicroServServer
}

func (s *server) DoJob(ctx context.Context, in *pb.ServRequest) (*pb.ServReply, error) {
	log.Printf("Label: %v, Params: %v", in.GetLabel(), in.GetParams())
	return &pb.ServReply{Response: "done job by serv-2."}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMicroServServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
