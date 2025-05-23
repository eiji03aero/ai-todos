package services

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	userRepo        repositories.UserRepository
	userSessionRepo repositories.UserSessionRepository
}

func NewAuthService(
	userRepo repositories.UserRepository,
	userSessionRepo repositories.UserSessionRepository,
) *AuthService {
	return &AuthService{
		userRepo:        userRepo,
		userSessionRepo: userSessionRepo,
	}
}

func (s *AuthService) CreateSession(ctx context.Context, userID string, isRemembered bool) (*models.UserSession, error) {
	session := models.NewUserSession(userID, isRemembered)
	if err := s.userSessionRepo.CreateSession(ctx, session); err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	return session, nil
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

func (s *AuthService) Login(
	ctx context.Context,
	email,
	password string,
	rememberMe bool,
) (*models.User, *models.UserSession, error) {
	// Validate input
	if err := models.ValidateEmail(email); err != nil {
		return nil, nil, err
	}

	// Validate password format
	if err := models.ValidatePassword(password); err != nil {
		return nil, nil, err
	}

	// Find user by email
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, errors.New("user not found")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, nil, errors.New("invalid credentials")
	}

	// Create user session
	session := models.NewUserSession(user.ID, rememberMe)
	if err := s.userSessionRepo.CreateSession(ctx, session); err != nil {
		return nil, nil, fmt.Errorf("error creating session: %w", err)
	}

	return user, session, nil
}

// ValidateSession checks if a given session token is valid
func (s *AuthService) ValidateSession(ctx context.Context, sessionToken string) (*models.User, error) {
	// Find the session
	session, err := s.userSessionRepo.FindSessionByToken(ctx, sessionToken)
	if err != nil {
		return nil, fmt.Errorf("error finding session: %w", err)
	}
	if session == nil {
		return nil, errors.New("session not found")
	}

	// Check if session is expired
	if session.IsExpired() {
		// Optional: Delete expired session
		_ = s.userSessionRepo.DeleteSession(ctx, sessionToken)
		return nil, errors.New("session expired")
	}

	// Fetch the associated user
	user, err := s.userRepo.FindByID(ctx, session.UserID)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Update last activity timestamp
	session.LastActivity = time.Now()
	_ = s.userSessionRepo.UpdateSession(ctx, session)

	return user, nil
}
