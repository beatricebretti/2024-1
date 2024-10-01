package main

import (
	"context"
	"log"
	"net"

	pb "labgrupo13/proto/namenode"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct {
	pb.UnimplementedNameNodeServer
}

func (s *server) DistributeDecisions(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	log.Printf("Recibido de mercenario %s", req.GetMercenaryId())
	return &pb.DecisionResponse{Status: "Distributed"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNameNodeServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error al servir: %v", err)
	}
}
