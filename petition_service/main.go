package main

import (
	"fmt"
	"net"

	"github.com/catness812/e-petitions-project/petition_service/internal/config"
	"github.com/catness812/e-petitions-project/petition_service/internal/controller/rpc"
	"github.com/catness812/e-petitions-project/petition_service/internal/pb"
	"github.com/catness812/e-petitions-project/petition_service/internal/repository"
	"github.com/catness812/e-petitions-project/petition_service/internal/service"
	"github.com/catness812/e-petitions-project/petition_service/pkg/database/postgres"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
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
		slog.Error(err)
		panic(err)
	}

	s := grpc.NewServer()
	server := &rpc.Server{
		PetitionService: petitionSvc,
	}

	pb.RegisterPetitionServiceServer(s, server)

	slog.Infof("gRPC Server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		slog.Error(err)
		panic(err)
	}
}
