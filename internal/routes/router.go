package routes

import (
	"gin-rest/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	// router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")

	authGroup := v1.Group("auth")
	auth := new(controllers.AuthController)
	authGroup.POST("/login", auth.Login)

	userGroup := v1.Group("user")
	user := new(controllers.UserController)
	userGroup.GET("/:id", user.Retrieve)

	return router
}
