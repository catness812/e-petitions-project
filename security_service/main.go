package main

import (
	"github.com/catness812/e-petitions-project/security_service/internal/config"
	"github.com/catness812/e-petitions-project/security_service/internal/controller"
	models "github.com/catness812/e-petitions-project/security_service/internal/model"
	"github.com/catness812/e-petitions-project/security_service/internal/pb"
	security_repository2 "github.com/catness812/e-petitions-project/security_service/internal/repository"
	"github.com/catness812/e-petitions-project/security_service/internal/service"
	"github.com/catness812/e-petitions-project/security_service/pkg/database/postgres"
	"github.com/catness812/e-petitions-project/security_service/pkg/database/rabbitmq"
	"github.com/catness812/e-petitions-project/security_service/pkg/database/redis_repository"
	"github.com/gookit/slog"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"net"
	"os"

	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"
)

var (
	redisDB    *redis.Client
	sService   *security_service.SecurityService
	sRpcServer *security_controller.SecurityRpcServer
	rabbitCh   *amqp.Channel
	cfg        *config.Config
)

func main() {
	RunSecurityService()
}

func RunSecurityService() {
	cfg = config.LoadConfig()
	dbConnection()
	userRepo := security_repository2.NewUserRepository(postgres.Database)
	redisRepo := security_repository2.NewRedisRepository(redisDB)
	sService = security_service.NewSecurityService(userRepo, redisRepo)
	sRpcServer = security_controller.NewSecurityRpcServer(sService, rabbitCh, cfg)
	lis, err := net.Listen("tcp", "localhost:9002")
	if err != nil {
		slog.Fatalf("Failed to listen to security service on GRPC port 9002: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSecurityServiceServer(grpcServer, sRpcServer)

	slog.Println("Listening security on 9002")
	if err := grpcServer.Serve(lis); err != nil {
		slog.Fatalf("failed to serve security service on 9002: %v", err)
	}
}

func dbConnection() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Fatalf("Some error occured. Err: %s", err)
	}
	postgres.Connect()
	err = postgres.Database.AutoMigrate(&models.UserModel{})
	if err != nil {
		slog.Fatalf("failed to automigrate: %v", err)
	}
	redisDB = redis_repository.NewRedisDBConnection()
	rabbitCh = rabbitmq.ConnectAMQPDataBase(os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), cfg.Rabbit.Host, cfg.Rabbit.Port)
}
