package routes

import (
	"snippet-sharing/config"
	"snippet-sharing/internal/controllers"
	"snippet-sharing/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := config.GetConfig()
	whitelistedHosts := cfg.GetStringSlice("security.whitelisted_hosts")
	sslRedirects := cfg.GetBool("security.ssl_redirects")

	router.Use(middlewares.HostValidationMiddleware(whitelistedHosts))
	router.Use(middlewares.SSLRedirectMiddleware(sslRedirects))
	router.Use(middlewares.SecurityHeadersMiddleware())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")

	authGroup := v1.Group("auth")
	auth := new(controllers.AuthController)
	authGroup.POST("/login", auth.Login)

	userGroup := v1.Group("user")
	userGroup.Use(middlewares.AuthMiddleware())
	user := new(controllers.UserController)
	userGroup.GET("/:id", user.Retrieve)

	return router
}
