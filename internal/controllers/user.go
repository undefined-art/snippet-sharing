package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (h UserController) Retrieve(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID required"})
		return
	}

	// TODO: Implement actual user retrieval from database
	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"username": "user_" + id,
		"message":  "User retrieval not yet implemented",
	})
}
