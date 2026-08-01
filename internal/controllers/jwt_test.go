package controllers

import (
	"testing"
	"time"

	"snippet-sharing/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	secret := "test-secret"
	username := "testuser"

	tokenStr, err := GenerateToken(username, secret)
	require.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

	claims := &types.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	require.NoError(t, err)
	assert.True(t, token.Valid)
	assert.Equal(t, username, claims.Username)
	assert.WithinDuration(t, time.Now().Add(24*time.Hour), claims.ExpiresAt.Time, 5*time.Second)
}

func TestGenerateToken_DifferentSecrets(t *testing.T) {
	tokenStr, err := GenerateToken("user", "secret1")
	require.NoError(t, err)

	claims := &types.Claims{}
	_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret2"), nil
	})
	assert.Error(t, err)
}