package helpers

import (
	"task-5-pbi-btpns-charles/database"
	"task-5-pbi-btpns-charles/helpers"
	"task-5-pbi-btpns-charles/models"
	"task-5-pbi-btpns-charles/helpers"
	"github.com/dgrijalva/jwt"
	"golang.org/x/crypto/bcrypt"
)

const jwtSecret = "secret-key" // Ganti dengan kunci rahasia JWT Anda

// GenerateJWTToken digunakan untuk menghasilkan token JWT berdasarkan ID pengguna
func GenerateJWTToken(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
    })

    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// GetUserIDFromToken digunakan untuk mendapatkan ID pengguna dari token JWT
func GetUserIDFromToken(tokenString string) (uint, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(jwtSecret), nil
    })

    if err != nil {
        return 0, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, err
    }

    userID, ok := claims["user_id"].(float64)
    if !ok {
        return 0, err
    }

    return uint(userID), nil
}

// HashPassword digunakan untuk menghasilkan hash kata sandi
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// VerifyPassword digunakan untuk memverifikasi kata sandi yang dihash

func VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
