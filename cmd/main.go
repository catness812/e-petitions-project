package main

import (
	"fmt"
	"net"
	"user_service/config"
	"user_service/internal/controller"
	"user_service/internal/pb"
	"user_service/internal/repository"
	"user_service/internal/service"
	"user_service/pkg/postgres"

	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var DBPostgres *gorm.DB

func main() {
	DBPostgres = postgres.Connect()
	userRepo := repository.NewUserRepository(DBPostgres)
	userService := service.NewUserService(userRepo)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.GrpcPort))
	if err != nil {
		slog.Error("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	UserController := controller.NewUserController(userService)

	pb.RegisterUserControllerServer(grpcServer, UserController)

	slog.Info("gRPC server listening on port %d", config.Cfg.GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		slog.Error("Failed to serve: %v", err)
	}
}
