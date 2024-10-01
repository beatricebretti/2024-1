package main

import (
	"context"
	"log"
	"net"

	pb "labgrupo13/proto/datanode"

	"google.golang.org/grpc"
)

const (
	port = ":50053"
)

type server struct {
	pb.UnimplementedDataNodeServer
}

func (s *server) StoreDecision(ctx context.Context, req *pb.DecisionRequest) (*pb.DecisionResponse, error) {
	log.Printf("Guardando decision mercenario %s", req.GetMercenaryId())
	return &pb.DecisionResponse{Status: "Stored"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDataNodeServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error al servir: %v", err)
	}
}
