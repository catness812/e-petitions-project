package main

import (
	"log"

	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	r := gin.Default()

	// authsvc.RegisterRoutes(r, &c)
	// usersvc.RegisterRoutes(r, &c)
	// adminsvc.RegisterRoutes(r, &c)

	r.Run(":1337")

}
