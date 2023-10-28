package database

import (
	"task-5-pbi-btpns-charles/models"

	"gorm.io/driver/mysql" 
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "user:password@tcp(localhost:3306)/your-database" // Ganti sesuai dengan informasi database Anda
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    DB = db
}

func MigrateDB() {
    DB.AutoMigrate(&models.User{}, &models.Photo{}) // Migrate model-model ke database
}
