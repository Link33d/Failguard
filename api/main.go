package main

import (
	"server/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	routes.Initialize(server)

	server.Run(":8000")

}
