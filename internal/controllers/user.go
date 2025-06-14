package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (h UserController) Retrieve(c *gin.Context) {
	c.String(http.StatusOK, "Alive!")
}
