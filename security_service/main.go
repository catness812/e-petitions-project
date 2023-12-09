package main

import (
	"fmt"
	"github.com/catness812/e-petitions-project/security_service/internal/config"
	"github.com/catness812/e-petitions-project/security_service/internal/controller"
	"github.com/catness812/e-petitions-project/security_service/internal/pb"
	security_repository "github.com/catness812/e-petitions-project/security_service/internal/repository"
	"github.com/catness812/e-petitions-project/security_service/internal/service"
	"github.com/catness812/e-petitions-project/security_service/pkg/database/postgres"
	"github.com/catness812/e-petitions-project/security_service/pkg/database/rabbitmq"
	"github.com/catness812/e-petitions-project/security_service/pkg/database/redis_repository"
	"github.com/gookit/slog"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"net"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"
)

func main() {
	RunSecurityService()
}

func RunSecurityService() {
	cfg := config.LoadConfig()
	redisDB, rabbitCh, postgresDB := dbConnection(cfg)
	userRepo := security_repository.NewUserRepository(postgresDB)
	redisRepo := security_repository.NewRedisRepository(redisDB)
	sService := security_service.NewSecurityService(userRepo, redisRepo)
	sRpcServer := security_controller.NewSecurityRpcServer(sService, rabbitCh, cfg)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GrpcPort))
	if err != nil {
		slog.Fatalf("Failed to listen to security service on GRPC port %v: %v", cfg.GrpcPort, err)
	}

	srvMetrics := newServerMetrics(cfg)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(),
		))
	pb.RegisterSecurityServiceServer(grpcServer, sRpcServer)

	slog.Infof("Listening security on %v", cfg.GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		slog.Fatalf("failed to serve security service on %v: %v", cfg.GrpcPort, err)
	}
}

func dbConnection(cfg *config.Config) (*redis.Client, *amqp.Channel, *gorm.DB) {
	postgresDB := postgres.Connect()
	rabbitUser, rabbitPass := fetchEnvVariables()
	redisDB := redis_repository.NewRedisDBConnection()
	rabbitCh := rabbitmq.ConnectAMQPDataBase(rabbitUser, rabbitPass, cfg.Rabbit.Host, cfg.Rabbit.Port)
	return redisDB, rabbitCh, postgresDB
}

func fetchEnvVariables() (string, string) {
	if err := godotenv.Load(".env"); err != nil {
		slog.Fatalf("Failed to read .env file: %v", err)
	}
	return os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS")
}

// newServerMetrics initializes Prometheus metrics with gRPC method interceptors
// and sets up an HTTP endpoint for Prometheus to collect the metrics
func newServerMetrics(cfg *config.Config) *grpcprom.ServerMetrics {
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.HttpPort), nil); err != nil {
			slog.Fatal(err)
		}
	}()
	return srvMetrics
}
