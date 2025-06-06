package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "devlink-backend/internal/db"
    "devlink-backend/internal/models"
)

func GetResources(c *gin.Context) {
    var resources []models.Resource

    roleVal, exists := c.Get("role")
    if exists && roleVal.(string) == "admin" {
       
        db.DB.Find(&resources)
    } else {
        
        db.DB.Where("approved = ?", true).Find(&resources)
    }

    c.JSON(http.StatusOK, resources)
}


func GetPendingResources(c *gin.Context) {
    role := c.MustGet("role").(string)
    if role != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
        return
    }

    var resources []models.Resource
    db.DB.Where("approved = ?", false).Find(&resources)
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

    
    if err := db.DB.Create(&resource).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit resource"})
        return
    }

    
    c.JSON(http.StatusOK, gin.H{
        "message":    "submitted for review",
        "resourceID": resource.ID, 
    })
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
