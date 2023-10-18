package user

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

func RegisterUserRoutes(r *gin.Engine, cfg *config.Config, rbacCfg *config.PermissionsConfig) {
	svc := InitUserServiceClient(cfg)
	securityClient, err := security.InitAuthServiceClient(cfg)
	if err != nil {
		slog.Fatalf("Failed to connect to security service : %v", err)
	}
	userrepo, err := NewUserRepository(cfg, svc)
	if err != nil {
		slog.Fatalf("Failed to connect to user repository : %v", err)
	}
	usersvc, err := NewUserService(userrepo)

	if err != nil {
		slog.Fatalf("Failed to connect to user service : %v", err)
	}

	userctrl := NewUserController(usersvc)
	authorizeMiddleware := middleware.NewAuthorizationMiddleware(svc, rbacCfg)
	authenticateMiddleware := middleware.NewAuthenticationMiddleware(securityClient)
	route := r.Group("/user")
	route.POST("", userctrl.CreateUser)
	route.GET("", authenticateMiddleware.Auth(), userctrl.GetUserByEmail)
	route.GET("/:uid", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("read", "user"), userctrl.GetUserByID)
	route.POST("/update", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("write", "user"), userctrl.UpdateUser)
	route.DELETE("", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("delete", "user"), userctrl.DeleteUser)
	route.POST("/admin", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("write", "user"), userctrl.AddAdmin)
}
