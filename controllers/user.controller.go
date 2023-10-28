package controllers

import (
	"net/http"
	"strconv"
	"task-5-pbi-btpns-charles/database"
	"task-5-pbi-btpns-charles/helpers"
	"task-5-pbi-btpns-charles/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if _, err := govalidator.ValidateStruct(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Lakukan validasi unik (contoh: email harus unik)
    if exists, _ := database.CheckIfUserExists(user.Email); exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah terdaftar"})
        return
    }

    user.Password = helpers.HashPassword(user.Password)

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Buat dan kirim token JWT setelah registrasi berhasil
    token, err := helpers.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"token": token})
}

func LoginUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Email tidak terdaftar"})
        return
    }

    // Verifikasi password (gunakan bcrypt atau metode keamanan lainnya)
    if !helpers.VerifyPassword(user.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Kata sandi salah"})
        return
    }

    // Buat dan kirim token JWT setelah login berhasil
    token, err := helpers.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
    // Mendapatkan ID pengguna dari JWT token
    userID, err := helpers.GetUserIDFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Mendapatkan ID pengguna yang akan diperbarui
    targetUserID, err := strconv.Atoi(c.Param("userId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
        return
    }

    // Hanya pengguna yang sesuai dengan ID JWT yang dapat diperbarui
    if userID != uint(targetUserID) {
        c.JSON(http.StatusForbidden, gin.H{"error": "Tidak diizinkan untuk memperbarui pengguna lain"})
        return
    }

    var user models.User
    if err := database.DB.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
        return
    }

    // Mendapatkan data pembaruan dari body JSON request
    var updateUser models.User
    if err := c.ShouldBindJSON(&updateUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Perbarui atribut-atribut yang diizinkan (contoh: username dan email)
    user.Username = updateUser.Username
    user.Email = updateUser.Email

    // Perbarui kata sandi jika diberikan
    if updateUser.Password != "" {
        // Implementasi hash dan penyimpanan kata sandi sesuai kebutuhan Anda
        user.Password = helpers.HashPassword(updateUser.Password)
    }

    if err := database.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profil pengguna berhasil diperbarui"})
}

func DeleteUser(c *gin.Context) {
    // Mendapatkan ID pengguna dari JWT token
    userID, err := helpers.GetUserIDFromToken(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Mendapatkan ID pengguna yang akan dihapus
    targetUserID, err := strconv.Atoi(c.Param("userId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
        return
    }

    // Hanya pengguna yang sesuai dengan ID JWT yang dapat dihapus
    if userID != uint(targetUserID) {
        c.JSON(http.StatusForbidden, gin.H{"error": "Tidak diizinkan untuk menghapus pengguna lain"})
        return
    }

    var user models.User
    if err := database.DB.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
        return
    }

    if err := database.DB.Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profil pengguna berhasil dihapus"})
}

