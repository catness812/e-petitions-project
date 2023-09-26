package security

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterSecurityRoutes(r *gin.Engine, c config.Config) {
	securitysvc, err := NewSecurityService(c)

	if err != nil {
		log.Fatal("Failed to connect to security service grpc: ", err)
	}

	userctrl := NewSecurityController(securitysvc)
	r.POST("/login", userctrl.Login)
	r.GET("/refresh", userctrl.Refresh)

}
