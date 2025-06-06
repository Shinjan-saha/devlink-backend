package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "devlink-backend/internal/models"
    "devlink-backend/internal/db"
    "devlink-backend/internal/auth"

	"os"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

   
    var existingUser models.User
    if err := db.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
        
        c.JSON(http.StatusConflict, gin.H{"error": "username or email already taken"})
        return
    }

    
    role := "user"
    if user.Role == "admin" {
        adminSecret := c.GetHeader("X-Admin-Secret")
        if adminSecret == "" {
            adminSecret = c.Query("admin_secret")
        }

        if adminSecret != os.Getenv("ADMIN_SECRET") {
            c.JSON(http.StatusForbidden, gin.H{"error": "not authorized to register as admin"})
            return
        }
        role = "admin"
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    user.Password = string(hashedPassword)
    user.Role = role

    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "registered", "role": user.Role})
}


func Login(c *gin.Context) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    db.DB.Where("email = ?", credentials.Email).First(&user)

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    token, _ := auth.GenerateToken(user.ID, user.Role)
    c.JSON(http.StatusOK, gin.H{"token": token})
}
