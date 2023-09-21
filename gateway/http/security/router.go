package security

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/gin-gonic/gin"
)

func RegisterSecurityRoutes(r *gin.Engine, c config.Config) {
	securitysvc, err := NewSecurityService(c)

	if err != nil {
		panic(err)
	}

	userctrl := NewSecurityController(securitysvc)
	r.POST("/login", userctrl.Login)
	r.GET("/refresh", userctrl.Refresh)

}
