package user

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterUserRoutes(r *gin.Engine, cfg *config.Config, rbacCfg *config.PermissionsConfig) {
	svc := InitUserServiceClient(cfg)
	securityClient := security.InitAuthServiceClient(cfg)
	userrepo, err := NewUserRepository(cfg, svc)
	if err != nil {
		log.Fatalf("Failed to connect to user service grpc: %v", err)
	}
	usersvc, err := NewUserService(userrepo)

	userctrl := NewUserController(usersvc)
	authorizeMiddleware := middleware.NewAuthorizationMiddleware(svc, rbacCfg)
	authenticateMiddleware := middleware.NewAuthenticationMiddleware(securityClient)
	route := r.Group("/user")
	route.POST("/", userctrl.CreateUser)
	route.GET("/", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("read", "user"), userctrl.GetUser)
	route.POST("/update", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("write", "user"), userctrl.UpdateUser)
	route.DELETE("/", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("delete", "user"), authorizeMiddleware.Authorize("delete", "users"), userctrl.DeleteUser)
	route.POST("/admin", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("write", "user"), userctrl.AddAdmin)
}
