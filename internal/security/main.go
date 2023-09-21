package main

import (
	"log"
	"net"

	models "github.com/catness812/e-petitions-project/internal/model"
	security_controller "github.com/catness812/e-petitions-project/internal/security/controller"
	"github.com/catness812/e-petitions-project/internal/security/pb"
	"github.com/catness812/e-petitions-project/internal/security/repository"
	security_service "github.com/catness812/e-petitions-project/internal/security/service"
	"github.com/catness812/e-petitions-project/pkg/database/postgres"
	"github.com/catness812/e-petitions-project/pkg/database/redis_repository"
	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"
)

var (
	redisDB    *redis.Client
	sService   *security_service.SecurityService
	sRpcServer *security_controller.SecurityRpcServer
)

func init() {
	postgres.Connect()
	err := postgres.Database.AutoMigrate(&models.UserModel{})
	if err != nil {
		log.Fatalf("failed to automigrate: %v", err)
	}
	redisDB = redis_repository.NewRedisDBConnection()

	userRepo := repository.NewUserRepository(postgres.Database)
	redisRepo := repository.NewRedisRepository(redisDB)
	sService = security_service.NewSecurityService(userRepo, redisRepo)
	sRpcServer = security_controller.NewSecurityRpcServer(sService)
}

func main() {
	RunSecurityService()
}

func RunSecurityService() {
	lis, err := net.Listen("tcp", "localhost:9002")
	if err != nil {
		log.Fatalf("Failed to listen to security service on GRPC port 9002: %v", err)
	}
	grpcServer := grpc.NewServer()
	//security_pb.RegisterRpcServiceServer(grpcServer, sRpcServer)
	//pb.RegisterSecurityServiceServer(grpcServer, security_controller.NewSecurityRpcServer(grpcServer, sRpcServer))

	pb.RegisterSecurityServiceServer(grpcServer, sRpcServer)

	log.Println("Listening security on 9002")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve security service on 9002: %v", err)
	}
}