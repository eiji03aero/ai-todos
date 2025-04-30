package services

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/eijiosakabe/ai-todos/api/internal/domain/models"
	"github.com/eijiosakabe/ai-todos/api/internal/domain/repositories"
)

var (
	ErrEmailExists     = errors.New("email already exists")
	ErrInvalidPassword = errors.New("invalid password")
)

// ValidationError represents a detailed validation error
type ValidationError struct {
	Field   string
	Message string
	Errors  []error
}

// Error implements the error interface
func (ve *ValidationError) Error() string {
	errorMessages := make([]string, len(ve.Errors))
	for i, err := range ve.Errors {
		errorMessages[i] = err.Error()
	}
	return fmt.Sprintf("%s: %v", ve.Message, errorMessages)
}

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) SignUp(ctx context.Context, email, password, passwordConfirmation string) (*models.User, error) {
	// Validate input
	validationErrors := models.ValidateUserSignup(email, password, passwordConfirmation)
	if len(validationErrors) > 0 {
		// Return all validation errors
		return nil, &ValidationError{
			Field:   "signup",
			Message: "Multiple validation errors occurred",
			Errors:  validationErrors,
		}
	}

	// Check email uniqueness
	exists, err := s.userRepo.EmailExists(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error checking email uniqueness: %w", err)
	}
	if exists {
		return nil, &ValidationError{
			Field:   "email",
			Message: "Email already exists",
		}
	}

	// Hash password securely
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	// Create user
	user := models.NewUser(email, string(hashedPassword))
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return user, nil
}

func (s *AuthService) CheckEmailUniqueness(ctx context.Context, email string) (bool, error) {
	// Validate email format
	if err := models.ValidateEmail(email); err != nil {
		return false, err
	}

	// Check email uniqueness
	exists, err := s.userRepo.EmailExists(ctx, email)
	if err != nil {
		return false, err
	}
	return !exists, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.User, error) {
	// Find user by email
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
