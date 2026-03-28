package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// UpdateLocation will eventually save coordinates to the database
func UpdateLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Location received!",
	})
}