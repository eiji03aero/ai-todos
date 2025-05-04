package repositories

import (
	"context"
	"errors"

	"github.com/eijiosakabe/ai-todos/api/internal/domain/models"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type UserRepository interface {
	// Check if an email already exists in the database
	EmailExists(ctx context.Context, email string) (bool, error)

	// Create a new user in the database
	CreateUser(ctx context.Context, user *models.User) error

	// Find a user by email
	FindByEmail(ctx context.Context, email string) (*models.User, error)

	// Find a user by ID
	FindByID(ctx context.Context, id string) (*models.User, error)
}
