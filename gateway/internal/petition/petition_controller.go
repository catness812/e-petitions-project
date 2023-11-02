package petition

import (
	"github.com/gofiber/fiber/v2"
	"strconv"

	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gookit/slog"
)

type IPetitionController interface {
	CreatePetition(ctx *fiber.Ctx) error
	GetPetitionByID(ctx *fiber.Ctx) error
	GetPetitions(ctx *fiber.Ctx) error
	UpdatePetitionStatus(ctx *fiber.Ctx) error
	DeletePetition(ctx *fiber.Ctx) error
	ValidatePetitionID(ctx *fiber.Ctx) error
	CreateVote(ctx *fiber.Ctx) error
	UpdatePetition(ctx *fiber.Ctx) error
	GetUserPetitions(ctx *fiber.Ctx) error
	GetUserVotedPetitions(ctx *fiber.Ctx) error
	GetAllSimilarPetitions(ctx *fiber.Ctx) error
	SearchPetitionsByTitle(ctx *fiber.Ctx) error
}

type petitionController struct {
	service IPetitionService
}

func NewPetitionController(service IPetitionService) IPetitionController {
	return &petitionController{
		service: service,
	}
}

func (c *petitionController) CreatePetition(ctx *fiber.Ctx) error {
	var petition model.CreatePetition

	if err := ctx.BodyParser(&petition); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	resp, err := c.service.CreatePetition(petition)
	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	return ctx.Status(fiber.StatusCreated).JSON(resp)

}

func (c *petitionController) GetPetitionByID(ctx *fiber.Ctx) error {
	pid, err := strconv.ParseUint(ctx.Params("pid"), 10, 32)
	if err != nil {

		slog.Errorf("Failed to get the id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	id := uint32(pid)

	petition, err := c.service.GetPetitionByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Info("Petition retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(petition)
}

func (c *petitionController) GetPetitions(ctx *fiber.Ctx) error {
	pageStr := ctx.Params("page")
	limitStr := ctx.Params("limit")

	page, err := strconv.ParseUint(pageStr, 10, 32)
	if err != nil {

		slog.Errorf("Failed to get the page: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil {

		slog.Errorf("Failed to get the limit: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	petitions, err := c.service.GetPetitions(uint32(page), uint32(limit))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	slog.Infof("All petitions retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(petitions)

}

func (c *petitionController) UpdatePetitionStatus(ctx *fiber.Ctx) error {
	var status model.Status
	err := ctx.BodyParser(&status)
	if err != nil {

		slog.Errorf("Failed to bind status: %s", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	err = c.service.UpdatePetitionStatus(status.ID, status.Status)

	if err != nil {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	slog.Infof("Petition status updated successfully")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Petition status updated successfully"})
}

func (c *petitionController) UpdatePetition(ctx *fiber.Ctx) error {
	var petition model.UpdatePetition
	err := ctx.BodyParser(&petition)
	if err != nil {
		slog.Errorf("Failed to bind petition: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	err = c.service.UpdatePetition(petition)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Info("Petition created successfully")
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Petition updated successfully"})
}

func (c *petitionController) DeletePetition(ctx *fiber.Ctx) error {
	idParam := ctx.Params("pid")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {

		slog.Errorf("Failed to get the id: %s", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	idUint32 := uint32(id)

	err = c.service.DeletePetition(idUint32)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	slog.Infof("Petition deleted successfully")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Petition deleted successfully"})
}

func (c *petitionController) ValidatePetitionID(ctx *fiber.Ctx) error {
	var pid struct {
		PetitionID uint32 `json:"petition_id"`
	}

	if err := ctx.BodyParser(&pid); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	err := c.service.ValidatePetitionID(pid.PetitionID)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	return ctx.Status(fiber.StatusOK).SendString("Petition validation is successful")

}

func (c *petitionController) CreateVote(ctx *fiber.Ctx) error {
	uid, err := strconv.ParseUint(ctx.Params("uid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the user id: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	pid, err := strconv.ParseUint(ctx.Params("pid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the petition id: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	err = c.service.CreateVote(uint32(uid), uint32(pid))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Infof("Petition signed successfully")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Vote created successfully"})
}
func (c *petitionController) GetUserPetitions(ctx *fiber.Ctx) error {
	uid, err := strconv.ParseUint(ctx.Params("uid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the user id: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	page, err := strconv.ParseUint(ctx.Params("page"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the page: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	limit, err := strconv.ParseUint(ctx.Params("limit"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the limit: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := c.service.GetUserPetitions(uint32(uid), uint32(page), uint32(limit))

	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Infof("User created petitions retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(res)

}

func (c *petitionController) GetUserVotedPetitions(ctx *fiber.Ctx) error {
	uid, err := strconv.ParseUint(ctx.Params("uid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the user id: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	page, err := strconv.ParseUint(ctx.Params("page"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the page: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	limit, err := strconv.ParseUint(ctx.Params("limit"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the limit: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := c.service.GetUserVotedPetitions(uint32(uid), uint32(page), uint32(limit))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Infof("User voted petitions retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *petitionController) GetAllSimilarPetitions(ctx *fiber.Ctx) error {
	var title model.Petition
	err := ctx.BodyParser(&title)
	if err != nil {
		slog.Errorf("Failed to bind title: ", err)

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := c.service.GetAllSimilarPetitions(title.Title)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Infof("Similar petition retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *petitionController) SearchPetitionsByTitle(ctx *fiber.Ctx) error {

	var title model.Petition
	err := ctx.BodyParser(&title)
	if err != nil {

		slog.Errorf("Failed to bind title: %s", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	page, err := strconv.ParseUint(ctx.Params("page"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the page: %s", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	limit, err := strconv.ParseUint(ctx.Params("limit"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the limit: %s", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := c.service.SearchPetitionsByTitle(title.Title, uint32(page), uint32(limit))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	slog.Infof("Search by title successfully")
	return ctx.JSON(res)

}
