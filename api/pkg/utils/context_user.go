package utils

import (
	"errors"

	"github.com/eijiosakabe/ai-todos/api/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// GetCurrentUser retrieves the authenticated user from the Gin context
// Returns an error if the user is not found or is of an invalid type
func GetCurrentUser(c *gin.Context) (*models.User, error) {
	// Retrieve the user from the context set by AuthMiddleware
	userValue, exists := c.Get("currentUser")
	if !exists {
		return nil, errors.New("user not found in context")
	}

	// Type assert to the user model
	user, ok := userValue.(*models.User)
	if !ok {
		return nil, errors.New("invalid user type in context")
	}

	return user, nil
}
