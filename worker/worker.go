package main

import (
	"context"
	"log"

	commons "Desktop/mr/commons"
	pb "Desktop/mr/proto"
	"fmt"
	"os"
	"plugin"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run worker.go <plugin.so>")
		return
	}
	workerID := fmt.Sprintf("worker-%d", os.Getpid())

	p, err := plugin.Open(args[0])
	if err != nil {
		panic(err)
	}
	mapFunc, reduceFunc := commons.FindFuncs(p)

	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCoordinatorServiceClient(conn)

	resp, err := client.JobStatus(context.Background(), &pb.JobStatusRequest{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Job status: %v", resp.Done)
	for {
		task, err := client.GetTask(context.Background(), &pb.GetTaskRequest{})
		if err != nil {
			log.Fatalf("failed to get task: %v", err)
		}
		log.Printf("Received task: %v", task)
		if task == nil {
			break
		}

		switch task.TaskType {
		case pb.TaskType_TASK_TYPE_MAP:
			reduce_path := mapFunc(task.InputPath)
			client.ReportMapDone(context.Background(), &pb.ReportMapDoneRequest{
				InputPath:   task.InputPath,
				WorkerId:    workerID,
				ReducePaths: []string{reduce_path},
			})

		case pb.TaskType_TASK_TYPE_REDUCE:
			output_path := reduceFunc(task.InputPath)
			client.ReportReduceDone(context.Background(), &pb.ReportReduceDoneRequest{
				OutputPath:  output_path,
				ReduceIndex: task.ReduceIndex,
				WorkerId:    workerID,
			})

		case pb.TaskType_TASK_TYPE_WAIT:
			time.Sleep(200 * time.Millisecond)

		case pb.TaskType_TASK_TYPE_EXIT:
			return

		}
	}

}
