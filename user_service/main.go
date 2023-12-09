package main

import (
	"fmt"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"

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

	srvMetrics := newServerMetrics()
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(),
		))

	UserController := controller.NewUserController(userService)

	pb.RegisterUserServiceServer(grpcServer, UserController)

	slog.Info("gRPC server listening on port %d", config.Cfg.GrpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %v", err)
	}

	return nil
}

// newServerMetrics initializes Prometheus metrics with gRPC method interceptors
// and sets up an HTTP endpoint for Prometheus to collect the metrics
func newServerMetrics() *grpcprom.ServerMetrics {
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%v", config.Cfg.HttpPort), nil); err != nil {
			slog.Fatal(err)
		}
	}()
	return srvMetrics
}
