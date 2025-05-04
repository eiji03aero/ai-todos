package handlers

import (
	"github.com/eijiosakabe/ai-todos/api/internal/presentation/middleware"
	"github.com/eijiosakabe/ai-todos/api/internal/services"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all API routes for the handlers package
func RegisterRoutes(r *gin.RouterGroup, authService *services.AuthService) {
	handler := NewAuthHandler(authService)

	apiRoutes := r.Group("/api")

	publicRoutes := apiRoutes.Group("")
	publicRoutes.POST("/auth/signup", handler.SignUp)
	publicRoutes.POST("/auth/login", handler.Login)
	publicRoutes.GET("/auth/check-email", handler.CheckEmail)

	authRoutes := apiRoutes.Group("")
	authMiddleware := middleware.AuthMiddleware(authService)
	authRoutes.Use(authMiddleware)
	authRoutes.POST("/auth/validate-session", handler.ValidateSession)
	authRoutes.GET("/auth/app_state", handler.AppState)
}
