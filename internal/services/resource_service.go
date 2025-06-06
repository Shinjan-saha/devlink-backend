package services

import (
    "errors"

    "gorm.io/gorm"

    "devlink-backend/internal/db"
    "devlink-backend/internal/models"
)

func GetAllApprovedResources() ([]models.Resource, error) {
    var resources []models.Resource
    err := db.DB.Where("approved = ?", true).Find(&resources).Error
    return resources, err
}

func SubmitNewResource(title, url, desc, rtype, tags string, userID uint) error {
    resource := models.Resource{
        Title:       title,
        URL:         url,
        Description: desc,
        Type:        rtype,
        Tags:        tags,
        SubmittedBy: userID,
        Approved:    false,
    }

    return db.DB.Create(&resource).Error
}

func ApproveResourceByID(resourceID string) error {
    var resource models.Resource
    if err := db.DB.First(&resource, resourceID).Error; err != nil {
        return err
    }

    if resource.Approved {
        return errors.New("already approved")
    }

    resource.Approved = true
    return db.DB.Save(&resource).Error
}
