package database

import (
	"github.com/your-username/findmygo/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is a global variable we can use in other files to talk to the database
var DB *gorm.DB

func InitDB() {
	// Connect to (or create) the database file
	database, err := gorm.Open(sqlite.Open("findmygo.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// This line automatically creates the User and Device tables for you
	database.AutoMigrate(&models.User{}, &models.Device{})

	DB = database
}