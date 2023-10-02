package main

import (
	"log"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()

	user.RegisterUserRoutes(r, &config.Cfg)
	//petition.RegisterPetitionRoutes(r, c)
	security.RegisterSecurityRoutes(r, &config.Cfg)

	err := r.Run(":1337")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
