package main

import (
	"log"
	"server/src/config"
	"server/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	server := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found. Continuing with system env vars.")
	}

	config.InitDatabase()

	routes.Initialize(server)

	server.Run(":8000")

}
