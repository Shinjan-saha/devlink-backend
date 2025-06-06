package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "devlink-backend/internal/db"
    "devlink-backend/internal/models"
)

func GetResources(c *gin.Context) {
    var resources []models.Resource
    db.DB.Where("approved = ?", true).Find(&resources)
    c.JSON(http.StatusOK, resources)
}

func SubmitResource(c *gin.Context) {
    var resource models.Resource
    if err := c.ShouldBindJSON(&resource); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    userID := c.MustGet("user_id").(uint)
    resource.SubmittedBy = userID
    resource.Approved = false
    db.DB.Create(&resource)
    c.JSON(http.StatusOK, gin.H{"message": "submitted for review"})
}

func ApproveResource(c *gin.Context) {
    userRole := c.MustGet("role").(string)
    if userRole != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
        return
    }

    id := c.Param("id")
    db.DB.Model(&models.Resource{}).Where("id = ?", id).Update("approved", true)
    c.JSON(http.StatusOK, gin.H{"message": "resource approved"})
}
