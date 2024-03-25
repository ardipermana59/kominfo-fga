package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key untuk penandatanganan token
var jwtSecret = []byte(os.Getenv("JWT_KEY"))

// GenerateJWTToken digunakan untuk membuat token JWT untuk pengguna dengan ID yang diberikan
func GenerateJWTToken(userID uint) (string, error) {
	// Set expiration time
	expirationTime := time.Now().Add(24 * time.Hour) // Contoh: token berlaku selama 24 jam

	// Create claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   string(userID),
	}

	fmt.Println("ini claim di crypto")
	fmt.Println(claims)

	fmt.Println("ini user id di crypto")
	fmt.Println(userID)

	fmt.Println("ini Subject id di crypto")
	fmt.Println(string(userID))
	// Create token dengan penandatanganan HMAC dan metode hash SHA-256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token dengan secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
