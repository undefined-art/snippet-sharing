package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HostValidationMiddleware(whitelistedHosts []string) gin.HandlerFunc {
	if len(whitelistedHosts) == 0 {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		host := c.Request.Host
		allowed := false

		for _, h := range whitelistedHosts {
			if strings.EqualFold(h, host) {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}

		c.Next()
	}
}

func SSLRedirectMiddleware(enabled bool) gin.HandlerFunc {
	if !enabled {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		proto := c.Request.Header.Get("X-Forwarded-Proto")

		if proto == "" {
			if c.Request.TLS != nil {
				c.Next()
				return
			}

			url := c.Request.URL
			url.Scheme = "https"
			url.Host = c.Request.Host

			c.Redirect(http.StatusMovedPermanently, url.String())
			c.Abort()
			return
		}

		if proto != "https" {
			url := c.Request.URL
			url.Scheme = "https"
			url.Host = c.Request.Host

			c.Redirect(http.StatusMovedPermanently, url.String())
			c.Abort()
			return
		}

		c.Next()
	}
}

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		c.Next()
	}
}
