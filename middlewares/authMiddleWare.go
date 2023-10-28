package middlewares

import (
    "github.com/gin-gonic/gin"
    "task-5-pbi-btpns-charles/database"
	"task-5-pbi-btpns-charles/helpers"
	"task-5-pbi-btpns-charles/models"
	"task-5-pbi-btpns-charles/helpers"
    "net/http"
)

// AuthMiddleware adalah middleware untuk otentikasi JWT
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token JWT tidak ditemukan"})
            c.Abort()
            return
        }

        userID, err := helpers.GetUserIDFromToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token JWT tidak valid"})
            c.Abort()
            return
        }

        c.Set("userID", userID)

        c.Next()
    }
}
