package routes

import (
    "github.com/gin-gonic/gin"
    "devlink-backend/internal/handlers"
    "devlink-backend/internal/middleware"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.POST("/register", handlers.Register)
        api.POST("/login", handlers.Login)

        resource := api.Group("/resources")
        resource.GET("", handlers.GetResources)

        resource.Use(middleware.AuthMiddleware())
        {
            resource.POST("", handlers.SubmitResource)
            resource.PUT("/:id/approve", handlers.ApproveResource)
        }
    }
}
