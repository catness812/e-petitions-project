package user

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitUserControllerClient(c),
	}

	route := r.Group("/user")
	route.GET("/", svc.Get)
	route.POST("/", svc.Create)
	route.POST("/update", svc.Update)
	route.DELETE("/", svc.Delete)
	return svc
}
