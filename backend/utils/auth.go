package utils

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username   string `json:"username"`
	Email      string `json:"emai"`
	Role       string `json:"role"`
	EthAccount string `json:"ethAccount"`
	jwt.RegisteredClaims
}

func GenerateTokens(username, email, role, ethAccount string) (string, string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, Claims{
		Username:   username,
		Role:       role,
		Email:      email,
		EthAccount: ethAccount,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}).SignedString(secret)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
	}).SignedString(secret)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization Header Required!"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("username", claims.Username)
			c.Set("role", claims.Role)
			c.Set("email", claims.Email)
			c.Set("ethAccount", claims.EthAccount)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token: " + err.Error()})
		}
	}
}
