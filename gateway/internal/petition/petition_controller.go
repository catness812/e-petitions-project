package petition

import (
	"net/http"
	"strconv"

	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

type IPetitionController interface {
	CreatePetition(ctx *gin.Context)
	UpdatePetition(ctx *gin.Context)
	GetPetitions(ctx *gin.Context)
	DeletePetition(ctx *gin.Context)
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
	var petition model.Petition
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

func (c *petitionController) UpdatePetition(ctx *gin.Context) {
	var petition model.Petition
	err := ctx.BindJSON(petition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := c.service.UpdatePetition(petition.PetitionId, string(petition.Status))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, msg)

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

	petitions, err := c.service.GetPetitions(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, petitions)

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
