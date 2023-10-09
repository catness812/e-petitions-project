package main

import (
	config "github.com/catness812/e-petitions-project/file_service/internal"
	"github.com/catness812/e-petitions-project/file_service/internal/controller"
	"github.com/catness812/e-petitions-project/file_service/internal/pb"
	"github.com/catness812/e-petitions-project/file_service/internal/repository"
	"github.com/catness812/e-petitions-project/file_service/internal/service"
	"github.com/catness812/e-petitions-project/file_service/pkg/db/postgres"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

func main() {

}

func RunFileService() {
	cfg := config.LoadConfig()
	postgresDB := dbConnection(cfg)
	fileRepo := repository.NewFileRepository(postgresDB)
	fileSvc := service.NewFileService(fileRepo)
	fileRPCServer := controller.NewFileRPCServer(fileSvc)

	lis, err := net.Listen("tcp", "localhost:9002")
	if err != nil {
		slog.Fatalf("Failed to listen to file service on GRPC port 9009: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, fileRPCServer)
	if err := grpcServer.Serve(lis); err != nil {
		slog.Fatalf("failed to serve security service on 9002: %v", err)
	}
	slog.Info("Listening file on 9009")
}

func dbConnection(cfg *config.Config) *gorm.DB {
	postgresDB := postgres.Connect()
	return postgresDB
}
