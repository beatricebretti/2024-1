package main

import (
	"context"
	pb "labgrupo13/proto/mercenarios"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	mercenaryID = "Mercenary1"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error al conectar: %v", err)
	}
	defer conn.Close()
	c := pb.NewDirectorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ReportReadiness(ctx, &pb.ReadinessRequest{MercenaryId: mercenaryID})
	if err != nil {
		log.Fatalf("error al reportar si esta listo: %v", err)
	}
	log.Printf("estado: %s", r.GetStatus())
}
