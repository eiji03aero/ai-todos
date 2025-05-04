package middleware

import (
	"net/http"
	"strings"

	"github.com/eijiosakabe/ai-todos/api/internal/services"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware returns a Gin middleware that protects routes by validating session ID from Authorization header.
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: missing session"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix
		sessionId := strings.TrimPrefix(authHeader, "Bearer ")
		if sessionId == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid session format"})
			c.Abort()
			return
		}

		user, err := authService.ValidateSession(c.Request.Context(), sessionId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid or expired session"})
			c.Abort()
			return
		}

		// Store user info in context for downstream handlers if needed
		c.Set("currentUser", user)

		c.Next()
	}
}
