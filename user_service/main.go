package main

import (
	"fmt"
	"net"

	"github.com/catness812/e-petitions-project/user_service/internal/config"
	"github.com/catness812/e-petitions-project/user_service/internal/controller"
	"github.com/catness812/e-petitions-project/user_service/internal/pb"
	"github.com/catness812/e-petitions-project/user_service/internal/repository"
	"github.com/catness812/e-petitions-project/user_service/internal/service"
	"github.com/catness812/e-petitions-project/user_service/pkg/postgres"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var DBPostgres *gorm.DB

func main() {
	DBPostgres = postgres.Connect()
	userRepo := repository.NewUserRepository(DBPostgres)
	userService := service.NewUserService(userRepo)

	if err := startGRPCServer(userService); err != nil {
		slog.Error("Failed to start gRPC server: %v", err)
	}
}

func startGRPCServer(userService *service.UserService) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.GrpcPort))
	if err != nil {
		return fmt.Errorf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	UserController := controller.NewUserController(userService)

	pb.RegisterUserControllerServer(grpcServer, UserController)

	slog.Info("gRPC server listening on port %d", config.Cfg.GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %v", err)
	}

	return nil
}
