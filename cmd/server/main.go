package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "devlink-backend/internal/db"
    "devlink-backend/internal/routes"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    db.InitDB()

    router := gin.Default()
    routes.RegisterRoutes(router)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Starting server on port %s", port)
    http.ListenAndServe(":"+port, router)
}
