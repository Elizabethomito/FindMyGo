package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/your-username/findmygo/internal/database"
	"github.com/your-username/findmygo/internal/models"
)

func Register(c *gin.Context) {
	var newUser models.User

	// 1. Get the JSON data sent by the user (Email/Password)
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	// 2. Save the user to the database
	// (Note: In a real app, we would hash the password first!)
	result := database.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// 3. Send back a success message
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}