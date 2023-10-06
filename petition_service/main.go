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
	petitionSvc := service.NewPetitionService(petitionRepo, publisherRepo)

	//publisherSvc := service.InitNotificationService(publisherRepo)
	grpcStart(petitionSvc)
}

func grpcStart(petitionSvc rpc.IPetitionService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.GrpcPort))
	if err != nil {
		slog.Error(err)
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterPetitionServiceServer(s, &rpc.Server{
		PetitionService: petitionSvc,
	})

	slog.Infof("gRPC Server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		slog.Error(err)
		panic(err)
	}
}
