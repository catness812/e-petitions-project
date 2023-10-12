package petition

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

func RegisterPetitionRoutes(r *gin.Engine, c *config.Config) {
	svc := InitPetitonServiceClient(c)
	petitionrepo, err := NewPetitionRepository(c, svc)
	petitionService, err := NewPetitionService(petitionrepo)
	if err != nil {
		slog.Fatalf("Failed to connect to petition service grpc: %v", err)

	}

	petitionController := NewPetitionController(petitionService)

	route := r.Group("/petition")
	route.POST("/", petitionController.CreatePetition)
	route.GET("/:id", petitionController.GetPetitionByID)
	route.GET("/all/:page/:limit", petitionController.GetPetitions)
	route.POST("/status/", petitionController.UpdatePetitionStatus)
	route.DELETE("/:id", petitionController.DeletePetition)
	//route.GET("/", petitionController.ValidatePetitionID)
	route.POST("/sign/:uid/:pid", petitionController.CreateVote)

	route = r.Group("/user")
	route.GET("/petitions/:uid/:page/:limit", petitionController.GetUserPetitions)
	route.GET("/voted/:uid/:page/:limit", petitionController.GetUserVotedPetitions)
}
