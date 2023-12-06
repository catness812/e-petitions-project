package user

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	pb2 "github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"github.com/catness812/e-petitions-project/gateway/internal/user/pb"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(r *fiber.App, rbacCfg *config.PermissionsConfig, userCtrl *UserController, userClient pb.UserServiceClient, securityClient pb2.SecurityServiceClient) {
	authorizeMiddleware := middleware.NewAuthorizationMiddleware(userClient, rbacCfg)
	authenticateMiddleware := middleware.NewAuthenticationMiddleware(securityClient)
	route := r.Group("/user")
	route.Post("", userCtrl.CreateUser)
	route.Post("/otp", userCtrl.OTPCreateUser)
	route.Get("", userCtrl.GetUserByEmail)
	route.Get("/:uid", userCtrl.GetUserByID)
	route.Post("/update", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("update", "user"), userCtrl.UpdateUser)
	route.Delete("", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("delete", "user"), userCtrl.DeleteUser)
	route.Post("/admin", authenticateMiddleware.Auth(), authorizeMiddleware.Authorize("write", "user"), userCtrl.AddAdmin)
}
