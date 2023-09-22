package petition

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPetitionController interface {
	CreatePetition(ctx *gin.Context)
	UpdatePetition(ctx *gin.Context)
	GetPetition(ctx *gin.Context)
	DeletePetition(ctx *gin.Context)
	GetAllPetitions(ctx *gin.Context)
	SignPetition(ctx *gin.Context)
	SetStatus(ctx *gin.Context)
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
	ctx.JSON(http.StatusOK, "nice")

}

func (c *petitionController) UpdatePetition(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")

}

func (c *petitionController) GetPetition(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")

}

func (c *petitionController) DeletePetition(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")
}

func (c *petitionController) GetAllPetitions(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")
}

func (c *petitionController) SignPetition(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")
}

func (c *petitionController) SetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nice")
}
