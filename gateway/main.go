package main

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	rbacCfg := config.LoadConfigRBAC()

	r := gin.Default()

	user.RegisterUserRoutes(r, cfg, rbacCfg)
	//petition.RegisterPetitionRoutes(r, cfg)
	security.RegisterSecurityRoutes(r, cfg)

	err := r.Run(":1337")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
