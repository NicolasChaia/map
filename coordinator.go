package main

import (
	pb "Desktop/mr/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoordinatorServiceServer
}

func (s *server) GetTask(context.Context, *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	// Minimal stub: worker should poll again while coordinator logic is being implemented.
	return &pb.GetTaskResponse{TaskType: pb.TaskType_TASK_TYPE_WAIT}, nil
}

func (s *server) ReportMapDone(context.Context, *pb.ReportMapDoneRequest) (*pb.Ack, error) {
	return &pb.Ack{Ok: true, Message: "map task acknowledged"}, nil
}

func (s *server) ReportReduceDone(context.Context, *pb.ReportReduceDoneRequest) (*pb.Ack, error) {
	return &pb.Ack{Ok: true, Message: "reduce task acknowledged"}, nil
}

func (s *server) JobStatus(context.Context, *pb.JobStatusRequest) (*pb.JobStatusResponse, error) {
	return &pb.JobStatusResponse{Done: false}, nil
}

func main() {
	// listen on a unix socket
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// start grpc server
	s := grpc.NewServer()

	// register the protobuf service with the server
	pb.RegisterCoordinatorServiceServer(s, &server{})

	// start serving requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
