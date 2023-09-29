package user

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterUserRoutes(r *gin.Engine, cfg *config.Config) {
	svc := InitUserServiceClient(cfg)
	userrepo, err := NewUserRepository(cfg, svc)
	if err != nil {
		log.Fatalf("Failed to connect to user service grpc: %v", err)
	}
	usersvc, err := NewUserService(userrepo)

	userctrl := NewUserController(usersvc)
	authMiddleware := middleware.NewAuthMiddleware(svc)
	route := r.Group("/user")
	route.POST("/", userctrl.CreateUser)
	route.GET("/:email", userctrl.GetUser)
	route.POST("/update", userctrl.UpdateUser)
	route.DELETE("/", authMiddleware.Authorize("delete", "users"), userctrl.DeleteUser)
	route.POST("/admin", userctrl.AddAdmin)
}
