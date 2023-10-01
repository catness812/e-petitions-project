package main

import (
	"github.com/catness812/e-petitions-project/security_service/internal/controller"
	models "github.com/catness812/e-petitions-project/security_service/internal/model"
	security_repository2 "github.com/catness812/e-petitions-project/security_service/internal/repository"
	"github.com/catness812/e-petitions-project/security_service/internal/security_pb"
	"github.com/catness812/e-petitions-project/security_service/internal/service"
	"github.com/catness812/e-petitions-project/security_service/pkg/database/postgres"
	"github.com/catness812/e-petitions-project/security_service/pkg/database/redis_repository"
	"github.com/gookit/slog"
	"net"

	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"
)

var (
	redisDB    *redis.Client
	sService   *security_service.SecurityService
	sRpcServer *security_controller.SecurityRpcServer
)

func main() {
	RunSecurityService()
}

func RunSecurityService() {
	postgres.Connect()
	err := postgres.Database.AutoMigrate(&models.UserModel{})
	if err != nil {
		slog.Fatalf("failed to automigrate: %v", err)
	}
	redisDB = redis_repository.NewRedisDBConnection()

	userRepo := security_repository2.NewUserRepository(postgres.Database)
	redisRepo := security_repository2.NewRedisRepository(redisDB)
	sService = security_service.NewSecurityService(userRepo, redisRepo)
	sRpcServer = security_controller.NewSecurityRpcServer(sService)
	lis, err := net.Listen("tcp", "localhost:9002")
	if err != nil {
		slog.Fatalf("Failed to listen to security service on GRPC port 9002: %v", err)
	}
	grpcServer := grpc.NewServer()
	security_pb.RegisterSecurityServiceServer(grpcServer, sRpcServer)

	slog.Println("Listening security on 9002")
	if err := grpcServer.Serve(lis); err != nil {
		slog.Fatalf("failed to serve security service on 9002: %v", err)
	}
}
