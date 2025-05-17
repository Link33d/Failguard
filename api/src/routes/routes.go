package routes

import (
	"server/src/controllers"

	"github.com/gin-gonic/gin"
)

func Initialize(server *gin.Engine) {

	server.GET("/", controllers.Index)

	server.POST("/check", controllers.CreateCheck)

}
