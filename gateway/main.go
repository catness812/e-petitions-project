package main

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/middleware"
	"github.com/catness812/e-petitions-project/gateway/internal/petition"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
)

func main() {
	r := fiber.New()

	prometheus := fiberprometheus.New("gateway")
	prometheus.RegisterAt(r, "/metrics")

	r.Use(prometheus.Middleware)
	r.Use(middleware.RateLimiterMiddleware())
	r.Use(middleware.CorsMiddleware())
	registerRoutes(r)
	err := r.Listen(":1337")
	if err != nil {
		slog.Fatalf("Failed to start server: %v", err)
	}
}

func registerRoutes(r *fiber.App) {
	cfg := config.LoadConfig()
	rbacCfg := config.LoadConfigRBAC()
	securityClient, err := security.InitAuthServiceClient(cfg)
	if err != nil {
		slog.Fatalf("Failed to connect to security service grpc: %v", err)
	}
	userClient := user.InitUserServiceClient(cfg)
	securityRepo := security.NewSecurityRepository(cfg, securityClient)
	securitySvc := security.NewSecurityService(securityRepo)
	securityCtrl := security.NewSecurityController(securitySvc, userClient)

	userRepo := user.NewUserRepository(cfg, userClient)
	userSvc := user.NewUserService(userRepo)
	userCtrl := user.NewUserController(userSvc)

	user.RegisterUserRoutes(r, rbacCfg, userCtrl, userClient, securityClient)
	petition.RegisterPetitionRoutes(r, cfg)
	security.RegisterSecurityRoutes(r, securityCtrl, securityClient)

}
