package main

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	rbacCfg := config.LoadConfigRBAC()

	r := gin.Default()
	r.Use(corsMiddleware())
	user.RegisterUserRoutes(r, cfg, rbacCfg)
	petition.RegisterPetitionRoutes(r, cfg)
	security.RegisterSecurityRoutes(r, cfg)

	err := r.Run(":1337")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,Access-Control-Allow-Origin")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
