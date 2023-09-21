package main

import (
	"fmt"
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
	loadDatabase()
	grpcStart()
}

func loadDatabase() {
	postgres.Connect()
	postgres.Database.AutoMigrate(&models.Petition{})
}

func grpcStart() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.GrpcPort))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterPetitionServiceServer(s, &rpctransport.Server{})

	log.Printf("gRPC Server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Panic(err)
	}
}
