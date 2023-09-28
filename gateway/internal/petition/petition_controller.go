package petition

import (
	"net/http"

	"github.com/catness812/e-petitions-project/gateway/model"
	"github.com/gin-gonic/gin"
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
	err := ctx.BindJSON(&query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	petitions, err := c.service.GetPetitions(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, petitions)

}

func (c *petitionController) DeletePetition(ctx *gin.Context) {
	var petition model.Petition
	err := ctx.BindJSON(petition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg, err := c.service.DeletePetition(petition.PetitionId)

	ctx.JSON(http.StatusOK, msg)
}
