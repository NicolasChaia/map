package main

import (
	"context"
	"log"

	pb "Desktop/mr/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runWorker() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCoordinatorServiceClient(conn)

	resp, err := client.JobStatus(context.Background(), &pb.JobStatusRequest{})
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("Job done: %v (map %d/%d, reduce %d/%d)", resp.Done, resp.MapCompleted, resp.MapTotal, resp.ReduceCompleted, resp.ReduceTotal)
}
