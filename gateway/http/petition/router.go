package petition

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/gin-gonic/gin"
)

func RegisterPetitionRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitPetitionServiceClient(c),
	}
	route := r.Group("/petition")
	route.POST("/", svc.NewPetition)
	route.POST("/update", svc.Update)
	route.DELETE("/", svc.Delete)
	route.POST("/sign", svc.Sign)
	route.GET("/all", svc.AllPetitions)
	route.GET("/", svc.Petition)
	route.POST("/status", svc.UpdateStatus)
}
