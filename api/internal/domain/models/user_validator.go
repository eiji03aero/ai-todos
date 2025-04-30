package models

import (
	"errors"
	"regexp"
	"unicode"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

// ValidationError represents a detailed validation error
type ValidationError struct {
	Field   string
	Message string
}

// ValidateEmail checks if the email is valid
func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword checks password complexity
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	var (
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}

	return nil
}

// ValidatePasswordConfirmation checks if password and confirmation match
func ValidatePasswordConfirmation(password, confirmation string) error {
	if password != confirmation {
		return errors.New("password and confirmation do not match")
	}
	return nil
}

// ValidateUserSignup performs comprehensive validation for user signup
func ValidateUserSignup(email, password, passwordConfirmation string) []error {
	var errors []error

	// Validate email
	if emailErr := ValidateEmail(email); emailErr != nil {
		errors = append(errors, emailErr)
	}

	// Validate password
	if passwordErr := ValidatePassword(password); passwordErr != nil {
		errors = append(errors, passwordErr)
	}

	// Validate password confirmation
	if confirmationErr := ValidatePasswordConfirmation(password, passwordConfirmation); confirmationErr != nil {
		errors = append(errors, confirmationErr)
	}

	return errors
}
