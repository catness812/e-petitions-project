package main

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c := config.LoadConfig()

	r := gin.Default()

	r.Run(":8080")

}
