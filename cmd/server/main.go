package main

import (
	"github.com/gin-gonic/gin"
	"github.com/your-username/findmygo/internal/database"
	"github.com/your-username/findmygo/internal/handlers" // Add this import
	"net/http"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "FindMyGo is ready!"})
	})

	// Add the Registration route
	r.POST("/register", handlers.Register)

	r.Run(":8080")
}