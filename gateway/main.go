package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//c := config.LoadConfig()

	r := gin.Default()

	r.Run(":1337")

}
