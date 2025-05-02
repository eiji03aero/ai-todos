package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSession struct {
	UserID       string
	Token        string
	ExpiresAt    time.Time
	IsRemembered bool
	LastActivity time.Time
}

func NewUserSession(userID string, isRemembered bool) *UserSession {
	now := time.Now()

	// Set expiration based on remember_me flag
	var expiresAt time.Time
	if isRemembered {
		// Long-lived session (30 days)
		expiresAt = now.AddDate(0, 0, 30)
	} else {
		// Short-lived session (24 hours)
		expiresAt = now.Add(24 * time.Hour)
	}

	return &UserSession{
		UserID:       userID,
		Token:        uuid.NewString(), // Generate unique token
		ExpiresAt:    expiresAt,
		IsRemembered: isRemembered,
		LastActivity: now,
	}
}

func (s *UserSession) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
