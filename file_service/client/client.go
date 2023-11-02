package main

import (
	"context"
	"fmt"
	"github.com/catness812/e-petitions-project/file_service/internal/pb"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
)

func main() {
	// Set up a connection to the gRPC server
	maxMsgSize := 1024 * 1024 * 1000 // 10MB, matching the server configuration
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMsgSize),
			grpc.MaxCallSendMsgSize(maxMsgSize),
		),
	}
	conn, err := grpc.Dial("localhost:50055", opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewFileServiceClient(conn)

	imageData, err := ioutil.ReadFile("/home/ktruedat/GolandProjects/e-petitions-project/file_service/client/JOTARO.jpg")
	if err != nil {
		slog.Fatalf("failed to read file: %v", err)
	}

	// Make a gRPC request
	response, err := client.UploadFile(context.Background(), &pb.FileRequest{FileData: imageData, Type: "mp4"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Response: %v\n", response)
}
