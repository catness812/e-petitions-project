package user

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterUserRoutes(r *gin.Engine, c config.Config) {
	usersvc, err := NewUserService(c)
	if err != nil {
		log.Fatal("Failed to connect to user service grpc: ", err)

	}

	userctrl := NewUserController(usersvc)

	route := r.Group("/user")
	route.GET("/", userctrl.GetUser)
	route.POST("/", userctrl.CreateUser)
	route.POST("/update", userctrl.UpdateUser)
	route.DELETE("/", userctrl.DeleteUser)
}
