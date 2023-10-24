package user

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	pb2 "github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, rbacCfg *config.PermissionsConfig, userCtrl *UserController, userClient pb.UserServiceClient, securityClient pb2.SecurityServiceClient) {
	authorizeMiddleware := middleware.NewAuthorizationMiddleware(userClient, rbacCfg)
	authenticateMiddleware := middleware.NewAuthenticationMiddleware(securityClient)
	route := r.Group("/user")
	route.POST("", userCtrl.CreateUser)
	route.POST("/otp", userCtrl.OTPCreateUser)
	route.GET("", authenticateMiddleware.Auth(), userCtrl.GetUserByEmail)
	route.GET("/:uid", authenticateMiddleware.Auth(), userCtrl.GetUserByID)
	route.POST("/update", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("update", "user"), userCtrl.UpdateUser)
	route.DELETE("", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("delete", "user"), userCtrl.DeleteUser)
	route.POST("/admin", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("write", "user"), userCtrl.AddAdmin)
}
