package lesson2

import (
	"os"
	"system_of_monitoring_statistics/services/auth/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user models.User, tokenTTL time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(tokenTTL).Unix()
	claims["iat"] = time.Now().Unix()

	// 6.3
	secretKey := os.Getenv("SECRET")

	// 6.3
	signedToken, errSignedToken := token.SignedString([]byte(secretKey))
	if errSignedToken != nil {
		return signedToken, errSignedToken
	}

	return signedToken, nil
}
