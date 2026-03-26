package models

import "time"

// User defines who owns the account
type User struct {
    ID        uint      `json:"id"`
    Email     string    `json:"email"`
    Password  string    `json:"-"` // We store the hash, but never "show" it in JSON
    CreatedAt time.Time `json:"created_at"`
}

// Device defines the hardware being tracked
type Device struct {
    ID          uint      `json:"id"`
    UserID      uint      `json:"user_id"`     // Links the device to its owner
    Name        string    `json:"name"`        // e.g., "Work Laptop"
    Latitude    float64   `json:"latitude"`    // GPS North/South
    Longitude   float64   `json:"longitude"`   // GPS East/West
    LastUpdated time.Time `json:"last_updated"`
}