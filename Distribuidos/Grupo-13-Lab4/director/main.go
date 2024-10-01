package main

import (
	"context"
	"log"
	"net"

	pb "labgrupo13/proto/director"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedDirectorServer
}

func (s *server) ReportReadiness(ctx context.Context, req *pb.ReadinessRequest) (*pb.ReadinessResponse, error) {
	log.Printf("Mercenario %s esta listo", req.GetMercenaryId())
	return &pb.ReadinessResponse{Status: "Acknowledged"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDirectorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("eror al servir: %v", err)
	}
}
