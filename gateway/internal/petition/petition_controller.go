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
	GetAllSimilarPetitions(ctx *gin.Context)
	SearchPetitionsByTitle(ctx *gin.Context)
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
		slog.Errorf("Failed to bind petition: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}

	resp, err := c.service.CreatePetition(petition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Info("Petition created successfully")
	ctx.JSON(http.StatusCreated, gin.H{"petition_id": resp, "error": false, "message": "Petition created successfully"})
}
func (c *petitionController) GetPetitionByID(ctx *gin.Context) {
	pid, err := strconv.ParseUint(ctx.Param("pid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the id: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get the id", "error": true})

	}
	id := uint32(pid)

	petition, err := c.service.GetPetitionByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve petition", "error": true})
		return
	}

	slog.Info("Petition retrieved successfully")
	ctx.JSON(http.StatusOK, gin.H{"petition": petition, "error": false, "message": "Petition retrieved successfully"})
}
func (c *petitionController) GetPetitions(ctx *gin.Context) {
	pageStr := ctx.Param("page")
	limitStr := ctx.Param("limit")

	page, err := strconv.ParseUint(pageStr, 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the page: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'page' parameter", "error": true})
		return
	}

	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the limit: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'limit' parameter", "error": true})
		return
	}

	petitions, err := c.service.GetPetitions(uint32(page), uint32(limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Infof("All petitions retrieved successfully")
	ctx.JSON(http.StatusOK, gin.H{"petitions": petitions, "error": false, "message": "All petitions retrieved successfully"})

}
func (c *petitionController) UpdatePetitionStatus(ctx *gin.Context) {
	var status model.Status
	err := ctx.BindJSON(&status)
	if err != nil {
		slog.Errorf("Failed to bind status: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}
	err = c.service.UpdatePetitionStatus(status.ID, status.Status)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Infof("Petition status updated successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Petition status updated successfully", "error": false})
}
func (c *petitionController) DeletePetition(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("pid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the id: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get the id", "error": true})

	}
	id := uint32(idParam)

	err = c.service.DeletePetition(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Infof("Petition deleted successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Petition deleted successfully", "error": false})
}
func (c *petitionController) ValidatePetitionID(ctx *gin.Context) {
	var pid struct {
		PetitionID uint32 `json:"petition_id"`
	}
	if err := ctx.BindJSON(&pid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}

	err := c.service.ValidatePetitionID(pid.PetitionID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
	}
	ctx.JSON(http.StatusOK, "Petition validation is successful")

}
func (c *petitionController) CreateVote(ctx *gin.Context) {
	uid, err := strconv.ParseUint(ctx.Param("uid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the user id: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'uid' parameter", "error": true})
		return
	}

	pid, err := strconv.ParseUint(ctx.Param("pid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the petition id: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'pid' parameter", "error": true})
		return
	}

	err = c.service.CreateVote(uint32(uid), uint32(pid))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Infof("Petition signed successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "Vote created successfully", "error": true})

}
func (c *petitionController) GetUserPetitions(ctx *gin.Context) {

	uid, err := strconv.ParseUint(ctx.Param("uid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the user id: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'uid' parameter", "error": true})
		return
	}

	page, err := strconv.ParseUint(ctx.Param("page"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the page: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'page' parameter", "error": true})
		return
	}

	limit, err := strconv.ParseUint(ctx.Param("limit"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the limit: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'limit' parameter", "error": true})
		return
	}

	res, err := c.service.GetUserPetitions(uint32(uid), uint32(page), uint32(limit))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
	}

	slog.Infof("User created petitions retrieved successfully")
	ctx.JSON(http.StatusOK, gin.H{"user_petitions": res, "error": false, "message": "User created petitions retrieved successfully"})

}
func (c *petitionController) GetUserVotedPetitions(ctx *gin.Context) {
	uid, err := strconv.ParseUint(ctx.Param("uid"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the user id: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'uid' parameter", "error": true})
		return
	}
	page, err := strconv.ParseUint(ctx.Param("page"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the page: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'page' parameter", "error": true})
		return
	}

	limit, err := strconv.ParseUint(ctx.Param("limit"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the limit: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'limit' parameter", "error": true})
		return
	}

	res, err := c.service.GetUserVotedPetitions(uint32(uid), uint32(page), uint32(limit))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
	}

	slog.Infof("User voted petitions retrieved successfully")
	ctx.JSON(http.StatusOK, gin.H{"user_voted_petitions": res, "error": false, "message": "User voted petitions retrieved successfully"})

}

func (c *petitionController) GetAllSimilarPetitions(ctx *gin.Context) {
	var title model.Petition
	err := ctx.BindJSON(&title)
	if err != nil {
		slog.Errorf("Failed to bind title: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}
	res, err := c.service.GetAllSimilarPetitions(title.Title)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Infof("Similar petition retrieved successfully")
	ctx.JSON(http.StatusOK, gin.H{"petitions": res, "error": false, "message": "Similar petition retrieved successfully"})

}

func (c *petitionController) SearchPetitionsByTitle(ctx *gin.Context) {
	var title model.Petition
	err := ctx.BindJSON(&title)
	if err != nil {
		slog.Errorf("Failed to bind title: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "error": true})
		return
	}

	page, err := strconv.ParseUint(ctx.Param("page"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the page: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	limit, err := strconv.ParseUint(ctx.Param("limit"), 10, 32)
	if err != nil {
		slog.Errorf("Failed to get the limit: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid 'limit' parameter", "error": true})
		return
	}

	res, err := c.service.SearchPetitionsByTitle(title.Title, uint32(page), uint32(limit))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "error": true})
		return
	}

	slog.Infof("Search by title successfully")
	ctx.JSON(http.StatusOK, gin.H{"petitions": res, "error": false, "message": "Search by title successfully"})

}
