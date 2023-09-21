package main

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/catness812/e-petitions-project/gateway/http/security"
	"github.com/catness812/e-petitions-project/gateway/http/user"
	"github.com/gin-gonic/gin"
)

func main() {
	c := config.LoadConfig()

	r := gin.Default()

	user.RegisterUserRoutes(r, c)
	security.RegisterSecurityRoutes(r, c)

	r.Run(":1337")

}
