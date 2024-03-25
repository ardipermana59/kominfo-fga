package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ardipermana59/mygram/pkg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Secret key untuk verifikasi token JWT
var jwtSecret = []byte(os.Getenv("JWT_KEY"))

// AuthMiddleware adalah middleware untuk memeriksa token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dapatkan token dari header Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		authArr := strings.Split(tokenString, " ")
		if len(authArr) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Unauthorized"})
			return
		}
		if authArr[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Unauthorized"})
			return
		}

		jwtToken := authArr[1]

		// Parse token JWT
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Unauthorized"})
			return
		}

		// Set userID dari token ke dalam konteks Gin
		claims, _ := token.Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)
		fmt.Println("Ini user id")
		fmt.Println(userID)

		fmt.Println("Ini claims")
		fmt.Println(claims)
		c.Set("userID", userID)

		c.Next()
	}
}
