package main

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	c := config.LoadConfig()

	r := gin.Default()

	user.RegisterUserRoutes(r, c)
	//petition.RegisterPetitionRoutes(r, c)
	security.RegisterSecurityRoutes(r, c)

	err := r.Run(":1337")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
