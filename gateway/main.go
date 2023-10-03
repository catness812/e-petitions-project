package main

import (
	"log"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	rbacCfg := config.LoadConfigRBAC()

	r := gin.Default()

	user.RegisterUserRoutes(r, &config.Cfg, rbacCfg)
	petition.RegisterPetitionRoutes(r, &config.Cfg)
	security.RegisterSecurityRoutes(r, &config.Cfg)

	err := r.Run(":1337")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
