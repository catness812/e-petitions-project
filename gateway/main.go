package main

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.LoadConfig()
	rbacCfg := config.LoadConfigRBAC()

	r := gin.Default()

	co := cors.DefaultConfig()
	co.AllowOrigins = []string{"http://localhost:5173"}
	co.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	co.AllowHeaders = []string{"Content-Type", "*"}
	co.AllowHeaders = []string{"Allow-Control-Allow-Origin", "*"}

	r.Use(cors.New(co))

	user.RegisterUserRoutes(r, &config.Cfg, rbacCfg)
	petition.RegisterPetitionRoutes(r, &config.Cfg)
	security.RegisterSecurityRoutes(r, &config.Cfg)

	err := r.Run(":1337")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
