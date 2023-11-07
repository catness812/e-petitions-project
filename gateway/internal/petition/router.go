package petition

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
)

func RegisterPetitionRoutes(r *fiber.App, c *config.Config) {
	svc := InitPetitonServiceClient(c)
	petitionrepo, err := NewPetitionRepository(c, svc)
	if err != nil {
		slog.Fatalf("Failed to connect to petition repository: %v", err)
	}
	petitionService, err := NewPetitionService(petitionrepo)
	if err != nil {
		slog.Fatalf("Failed to connect to petition service: %v", err)
	}

	petitionController := NewPetitionController(petitionService)

	route := r.Group("/petition")

	// route.Post("/update", petitionController.UpdatePetition)
	route.Post("", petitionController.CreatePetition)
	route.Get("/:pid", petitionController.GetPetitionByID)
	route.Get("/all/:page/:limit", petitionController.GetPetitions)
	route.Post("/status", petitionController.UpdatePetitionStatus)
	route.Delete("/:pid", petitionController.DeletePetition)

	//route.GET("/", petitionController.ValidatePetitionID)
	route.Post("/sign/:uid/:pid", petitionController.CreateVote)
	route.Post("/search/:page/:limit", petitionController.SearchPetitionsByTitle)
	route.Post("/similar", petitionController.GetAllSimilarPetitions)

	route = r.Group("/user")
	route.Get("/petitions/:uid/:page/:limit", petitionController.GetUserPetitions)
	route.Get("/voted/:uid/:page/:limit", petitionController.GetUserVotedPetitions)
}
