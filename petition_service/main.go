package main

import (
	"fmt"
	"github.com/catness812/e-petitions-project/petition_service/internal/config"
	"github.com/catness812/e-petitions-project/petition_service/internal/controller/rpc"
	"github.com/catness812/e-petitions-project/petition_service/internal/pb"
	"github.com/catness812/e-petitions-project/petition_service/internal/repository"
	"github.com/catness812/e-petitions-project/petition_service/internal/service"
	"github.com/catness812/e-petitions-project/petition_service/pkg/database/postgres"
	"github.com/gookit/slog"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {
	config.LoadConfig()
	db := postgres.LoadDatabase()
	petitionRepo := repository.NewPetitionRepository(db)
	publisherRepo := repository.NewPublisherRepository()
	userRepo := repository.NewUserRepository()
	elasticSearchRepo := repository.NewElasticRepository()
	petitionSvc := service.NewPetitionService(petitionRepo, publisherRepo, userRepo, elasticSearchRepo)

	grpcStart(petitionSvc)
}

func grpcStart(petitionSvc rpc.IPetitionService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Cfg.GrpcPort))
	if err != nil {
		slog.Fatal(err)
	}

	srvMetrics := newServerMetrics()
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(),
		))
	server := &rpc.Server{
		PetitionService: petitionSvc,
	}

	pb.RegisterPetitionServiceServer(s, server)

	slog.Infof("gRPC Server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		slog.Fatal(err)
	}
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
