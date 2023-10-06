package petition

import (
	"net/http"
	"strconv"

	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gookit/slog"

	"github.com/gin-gonic/gin"
)

type IPetitionController interface {
	CreatePetition(ctx *gin.Context)
	GetPetitionByID(ctx *gin.Context)
	GetPetitions(ctx *gin.Context)
	UpdatePetitionStatus(ctx *gin.Context)
	DeletePetition(ctx *gin.Context)
	ValidatePetitionID(ctx *gin.Context)
	CreateVote(ctx *gin.Context)
	GetUserPetitions(ctx *gin.Context)
	GetUserVotedPetitions(ctx *gin.Context)
}

type petitionController struct {
	service IPetitionService
}

func NewPetitionController(service IPetitionService) IPetitionController {

	return &petitionController{
		service: service,
	}
}
func (c *petitionController) CreatePetition(ctx *gin.Context) {
	var petition model.CreatePetition
	err := ctx.BindJSON(&petition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.service.CreatePetition(petition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
func (c *petitionController) GetPetitionByID(ctx *gin.Context) {
	var pid struct {
		PetitionID uint32 `json:"petition_id"`
	}

	if err := ctx.ShouldBindJSON(&pid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	petition, err := c.service.GetPetitionByID(pid.PetitionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve petition"})
		return
	}

	ctx.JSON(http.StatusOK, petition)
}
func (c *petitionController) GetPetitions(ctx *gin.Context) {
	query := model.PaginationQuery{}
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 32)
	if err != nil {
		slog.Printf("Failed to get page value: ", err)
	}
	query.Page = uint32(page)

	limit, err := strconv.ParseUint(ctx.Query("limit"), 10, 32)
	if err != nil {
		slog.Printf("Failed to get limit value: ", err)
	}
	query.Limit = uint32(limit)

	petitions, err := c.service.GetPetitions(uint32(page), uint32(limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, petitions)

}
func (c *petitionController) UpdatePetitionStatus(ctx *gin.Context) {
	var status struct {
		id    uint32 `json:"id"`
		title string `json:"title"`
	}
	err := ctx.BindJSON(&status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.service.UpdatePetitionStatus(status.id, status.title)
}
func (c *petitionController) DeletePetition(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		slog.Printf("Failed to get the id: ", err)
	}
	id := uint32(idParam)

	err = c.service.DeletePetition(id)

	ctx.Status(http.StatusOK)
}

func (c *petitionController) ValidatePetitionID(ctx *gin.Context) {
	var pid struct {
		PetitionID uint32 `json:"petition_id"`
	}
	if err := ctx.BindJSON(&pid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.ValidatePetitionID(pid.PetitionID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, "Petition validation is successful")

}
func (c *petitionController) CreateVote(ctx *gin.Context) {
	var ids struct {
		UserID     uint32 `json:"user_id"`
		PetitionID uint32 `json:"petition_id"`
	}

	if err := ctx.BindJSON(&ids); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.CreateVote(ids.UserID, ids.PetitionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, "Voted successfully")

}
func (c *petitionController) GetUserPetitions(ctx *gin.Context) {
	var uid struct {
		UserID uint32 `json:"user_id"`
	}
	if err := ctx.BindJSON(&uid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pageStr := ctx.Param("page")
	limitStr := ctx.Param("limit")

	// Convert the page and limit strings to uint32
	page, err := strconv.ParseUint(pageStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'page' parameter"})
		return
	}
	pag := uint32(page)

	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'limit' parameter"})
		return
	}

	lim := uint32(limit)

	res, err := c.service.GetUserPetitions(uid.UserID, pag, lim)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, res)

}
func (c *petitionController) GetUserVotedPetitions(ctx *gin.Context) {
	var uid struct {
		UserID uint32 `json:"user_id"`
	}
	if err := ctx.BindJSON(&uid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pageStr := ctx.Param("page")
	limitStr := ctx.Param("limit")

	// Convert the page and limit strings to uint32
	page, err := strconv.ParseUint(pageStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'page' parameter"})
		return
	}
	pag := uint32(page)

	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'limit' parameter"})
		return
	}

	lim := uint32(limit)

	res, err := c.service.GetUserVotedPetitions(uid.UserID, pag, lim)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, res)

}
