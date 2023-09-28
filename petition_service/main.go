package main

import (
	"fmt"
	"log"
	"net"

	"github.com/catness812/e-petitions-project/petition_service/internal/config"
	rpctransport "github.com/catness812/e-petitions-project/petition_service/internal/controller/rpc"
	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/pb"
	"github.com/catness812/e-petitions-project/petition_service/internal/repository"
	"github.com/catness812/e-petitions-project/petition_service/internal/service"
	"github.com/catness812/e-petitions-project/petition_service/pkg/database/postgres"
	"google.golang.org/grpc"

	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()
	db := loadDatabase()
	petitionRepo := repository.InitPetitionRepository(db)
	petitionSvc := service.InitPetitionService(petitionRepo)
	grpcStart(petitionSvc)
}

// TODO move all this stuff to db pkg
func loadDatabase() *gorm.DB {
	db := postgres.Connect()
	err := db.AutoMigrate(&models.Petition{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Status{})
	if err != nil {
		log.Fatal(err)
	}

	seedStatuses(db)

	return db
}

// seeds Statuses in case there are none in the sb. TODO Definitely move this out of here
func seedStatuses(db *gorm.DB) {
	var count int64
	db.Model(&models.Status{}).Count(&count)
	if count == 0 {
		for _, status := range models.StatusSeedData {
			db.Create(&status)
		}
	}
}

func grpcStart(petitionSvc rpctransport.IPetitionService) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.GrpcPort))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterPetitionServiceServer(s, &rpctransport.Server{
		PetitionService: petitionSvc,
	})

	log.Printf("gRPC Server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Panic(err)
	}
}
