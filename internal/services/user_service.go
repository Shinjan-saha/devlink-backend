package services

import (
    "errors"

    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"

    "devlink-backend/internal/db"
    "devlink-backend/internal/models"
    "devlink-backend/internal/auth"
)

func CreateUser(username, email, password string) error {
    var existing models.User
    if err := db.DB.Where("email = ?", email).First(&existing).Error; err == nil {
        return errors.New("user already exists")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return err
    }

    user := models.User{
        Username: username,
        Email:    email,
        Password: string(hashedPassword),
        Role:     "user",
    }

    return db.DB.Create(&user).Error
}

func AuthenticateUser(email, password string) (string, error) {
    var user models.User
    if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return "", errors.New("invalid credentials")
        }
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token, err := auth.GenerateToken(user.ID, user.Role)
    if err != nil {
        return "", err
    }

    return token, nil
}
