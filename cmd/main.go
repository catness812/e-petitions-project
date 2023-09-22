package main

import (
	"fmt"
	"github.com/catness812/e-petitions-project/internal/repository"
	"github.com/catness812/e-petitions-project/internal/service"
	"gorm.io/gorm"
	"log"
	"net"

	"github.com/catness812/e-petitions-project/config"
	rpctransport "github.com/catness812/e-petitions-project/internal/controller/rpc-transport"
	"github.com/catness812/e-petitions-project/internal/models"
	"github.com/catness812/e-petitions-project/internal/pb"
	"github.com/catness812/e-petitions-project/pkg/database/postgres"
	"google.golang.org/grpc"
)

func main() {
	db := loadDatabase()
	petitionRepo := repository.InitPetitionRepository(db)
	petitionSvc := service.InitPetitionService(petitionRepo)
	grpcStart(petitionSvc)
}

func loadDatabase() *gorm.DB {
	db := postgres.Connect()
	db.AutoMigrate(&models.Petition{})
	return db
}

func grpcStart(petitionSvc rpctransport.IPetitionSvc) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.GrpcPort))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterPetitionServiceServer(s, &rpctransport.Server{
		Svc: petitionSvc,
	})

	log.Printf("gRPC Server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Panic(err)
	}
}
