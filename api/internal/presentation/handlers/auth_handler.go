package handlers

import (
	"net/http"
	"time"

	"github.com/eijiosakabe/ai-todos/api/internal/services"
	"github.com/eijiosakabe/ai-todos/api/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Ensure ValidationError is imported
var _ = services.ValidationError{}

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type SignUpRequest struct {
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"remember_me"`
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.authService.SignUp(c.Request.Context(), req.Email, req.Password, req.PasswordConfirmation)
	if err != nil {
		switch v := err.(type) {
		case *services.ValidationError:
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   v.Message,
				"details": v.Errors,
			})
		case error:
			switch v {
			case services.ErrEmailExists:
				c.JSON(http.StatusConflict, gin.H{
					"error": "Email already exists",
				})
			case services.ErrInvalidPassword:
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid password",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
			}
		}
		return
	}

	// Create user session after signup
	session, err := h.authService.CreateSession(c.Request.Context(), user.ID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create session",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"createdAt": user.CreatedAt.Format(time.RFC3339),
		"sessionId": session.Token,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, session, err := h.authService.Login(
		c.Request.Context(),
		req.Email,
		req.Password,
		req.RememberMe,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
		"sessionId": session.Token,
		"session": gin.H{
			"expires_at":    session.ExpiresAt.Format(time.RFC3339),
			"is_remembered": session.IsRemembered,
		},
	})
}

func (h *AuthHandler) CheckEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is required",
		})
		return
	}

	isUnique, err := h.authService.CheckEmailUniqueness(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error checking email",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isUnique": isUnique,
	})
}

func (h *AuthHandler) ValidateSession(c *gin.Context) {
	user, err := utils.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session"})
		return
	}

	// Since user model has no expiresAt, we omit it in response
	c.JSON(http.StatusOK, gin.H{
		"userId": user.ID,
	})
}

func (h *AuthHandler) AppState(c *gin.Context) {
	user, err := utils.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId": user.ID,
		"email":  user.Email,
	})
}
