package petition

import (
	"log"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func RegisterPetitionRoutes(r *gin.Engine, c config.Config) {
	svc := InitPetitonServiceClient(c)
	petitionrepo, err := NewPetitionRepository(c, svc)
	petitionService, err := NewPetitionService(petitionrepo)
	if err != nil {
		log.Fatal("Failed to connect to petition service grpc: ", err)

	}

	petitionController := NewPetitionController(petitionService)

	route := r.Group("/petition")
	route.POST("/", petitionController.CreatePetition)
	route.DELETE("/:id", petitionController.DeletePetition)
	route.POST("/update", petitionController.UpdatePetition)
	route.GET("/all", petitionController.GetPetitions)

}
