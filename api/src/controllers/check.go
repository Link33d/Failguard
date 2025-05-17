package controllers

import (
	"net/http"
	"server/src/models"
	"time"

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

	now := time.Now().Format(time.RFC3339)
	check.CreatedAt = now
	check.UpdatedAt = now

	ctx.JSON(http.StatusOK, gin.H{"message": "Check successfully created!"})

}
