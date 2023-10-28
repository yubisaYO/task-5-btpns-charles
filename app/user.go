package models

import (
    "time"
)

type User struct {
    ID        uint      `json:"id" gorm:"primary_key"`
    Username  string    `json:"username" binding:"required"`
    Email     string    `json:"email" binding:"required" gorm:"unique"`
    Password  string    `json:"password" binding:"required" gorm:"size:60"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
