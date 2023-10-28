package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "task-5-pbi-btpns-charles/database"
	"task-5-pbi-btpns-charles/helpers"
	"task-5-pbi-btpns-charles/models"
)

func CreatePhoto(c *gin.Context) {
    var photo models.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validasi dan set atribut user ID menggunakan JWT
    userID, err := helpers.GetUserIDFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    photo.UserID = userID

    if err := database.DB.Create(&photo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Foto profil berhasil ditambahkan"})
}

func DeletePhoto(c *gin.Context) {
    photoID, err := strconv.Atoi(c.Param("photoId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID foto profil tidak valid"})
        return
    }

    // Validasi dan set atribut user ID menggunakan JWT
    userID, err := helpers.GetUserIDFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var photo models.Photo
    if err := database.DB.Where("id = ? AND user_id = ?", photoID, userID).First(&photo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Foto profil tidak ditemukan"})
        return
    }

    if err := database.DB.Delete(&photo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Foto profil berhasil dihapus"})
}
package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "your-app/database"
    "your-app/models"
    "your-app/helpers"
)

func CreatePhoto(c *gin.Context) {
    var photo models.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validasi dan set atribut user ID menggunakan JWT
    userID, err := helpers.GetUserIDFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    photo.UserID = userID

    if err := database.DB.Create(&photo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Foto profil berhasil ditambahkan"})
}

func DeletePhoto(c *gin.Context) {
    photoID, err := strconv.Atoi(c.Param("photoId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID foto profil tidak valid"})
        return
    }

    // Validasi dan set atribut user ID menggunakan JWT
    userID, err := helpers.GetUserIDFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var photo models.Photo
    if err := database.DB.Where("id = ? AND user_id = ?", photoID, userID).First(&photo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Foto profil tidak ditemukan"})
        return
    }

    if err := database.DB.Delete(&photo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Foto profil berhasil dihapus"})
}
