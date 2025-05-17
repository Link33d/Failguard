package controllers

import (
	"net/http"
	"server/src/models"
	"server/src/services"

	"github.com/gin-gonic/gin"
)

func CreateCheck(ctx *gin.Context) {

	var check models.Check
	if err := ctx.ShouldBindJSON(&check); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	if err := check.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := services.SaveCheck(&check); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save check"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Check successfully created!", "data": check})

}
