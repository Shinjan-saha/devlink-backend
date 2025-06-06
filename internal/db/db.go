package db

import (
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "devlink-backend/internal/models"
)

var DB *gorm.DB

func InitDB() {
    dsn := os.Getenv("DB_URL")
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    err = DB.AutoMigrate(&models.User{}, &models.Resource{})
    if err != nil {
        log.Fatalf("Failed to migrate models: %v", err)
    }
}
