package controllers

import (
	"net/http"
	"snippet-sharing/config"
	"snippet-sharing/internal/types"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct{}

func GenerateToken(username string, secret string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &types.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (a AuthController) Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})

		return
	}

	configuration := config.GetConfig()
	jwtSecret := configuration.GetString("http.auth.secret")

	if len(jwtSecret) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret not configured"})

		return
	}

	// TODO: Implement proper user authentication with password hashing
	// For now, accept any non-empty credentials (placeholder)
	if creds.Username == "" || creds.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})

		return
	}

	token, err := GenerateToken(creds.Username, jwtSecret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
