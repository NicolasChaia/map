package main

import (
	pb "Desktop/mr/proto"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	structs "Desktop/mr/structs"

	"google.golang.org/grpc"
)

type filesTaken struct {
	file      *structs.File
	workerID  string
	startedAt time.Time
}

type server struct {
	pb.UnimplementedCoordinatorServiceServer
	mu          sync.Mutex
	shelve      structs.Shelve
	mapTaken    map[string]filesTaken // key: input_path
	reduceTaken map[int32]filesTaken
	nReduce     int32
	timeout     time.Duration
}

func (s *server) GetTask(context.Context, *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.shelve.AllFilesFinished() {
		return &pb.GetTaskResponse{TaskType: pb.TaskType_TASK_TYPE_EXIT}, nil
	}

	// Verifico que haya terminado con todos los maps
	if !s.shelve.AllFilesMapped() {

		fmt.Println("Asignando tarea de map")
		f := s.shelve.GetNextFileMap()
		if f == nil {
			return &pb.GetTaskResponse{TaskType: pb.TaskType_TASK_TYPE_WAIT}, nil
		}

		s.mapTaken[f.GetPath()] = filesTaken{file: f, startedAt: time.Now()}
		return &pb.GetTaskResponse{TaskType: pb.TaskType_TASK_TYPE_MAP, InputPath: f.GetPath(), NReduce: s.nReduce}, nil
	}

	//Se que ya termino con los maps entonces paso al reduce

	f := s.shelve.GetNextFileReduce()
	fmt.Println("Asignando tarea de reduce")
	if f == nil {
		return &pb.GetTaskResponse{TaskType: pb.TaskType_TASK_TYPE_WAIT}, nil
	}
	reduceIndex := int32(0)
	s.reduceTaken[reduceIndex] = filesTaken{file: f, startedAt: time.Now()}
	return &pb.GetTaskResponse{TaskType: pb.TaskType_TASK_TYPE_REDUCE, InputPath: f.ReducePaths[reduceIndex], ReduceIndex: reduceIndex, NReduce: s.nReduce}, nil
}

func (s *server) ReportMapDone(_ context.Context, req *pb.ReportMapDoneRequest) (*pb.Ack, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	taken, ok := s.mapTaken[req.InputPath]
	if !ok || taken.file == nil {
		return &pb.Ack{Ok: false, Message: "unknown input_path"}, nil
	}

	fmt.Printf("Map task done for file: %s by worker: %s\n", req.InputPath, req.WorkerId)
	s.shelve.MarkMapFinished(taken.file)
	for _, reducePath := range req.ReducePaths {
		taken.file.AddReducePath(reducePath)
	}
	delete(s.mapTaken, req.InputPath)

	return &pb.Ack{Ok: true, Message: "map task acknowledged"}, nil
}

func (s *server) ReportReduceDone(_ context.Context, req *pb.ReportReduceDoneRequest) (*pb.Ack, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	taken, ok := s.reduceTaken[req.ReduceIndex]
	if !ok || taken.file == nil {
		return &pb.Ack{Ok: false, Message: "unknown reduce index"}, nil
	}

	s.shelve.MarkFileFinished(taken.file)
	delete(s.reduceTaken, req.ReduceIndex)

	return &pb.Ack{Ok: true, Message: "reduce task acknowledged"}, nil
}

func (s *server) JobStatus(context.Context, *pb.JobStatusRequest) (*pb.JobStatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return &pb.JobStatusResponse{Done: s.shelve.AllFilesFinished()}, nil
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Usage: coordinator.go <input_files>")
	}
	inputFiles := args[:len(args)-1]
	srv := &server{
		shelve:      structs.NewShelve(),
		mapTaken:    make(map[string]filesTaken),
		reduceTaken: make(map[int32]filesTaken),
		nReduce:     1,
		timeout:     10 * time.Second,
	}
	srv.shelve.AddFiles(inputFiles)
	log.Printf("Starting coordinator with input files: %v", inputFiles)

	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCoordinatorServiceServer(s, srv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
