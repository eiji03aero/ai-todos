package repomodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/eijiosakabe/ai-todos/api/internal/domain/models"
)

type UserSession struct {
	gorm.Model
	UserID       string    `gorm:"type:uuid;not null;column:user_id"`
	Token        string    `gorm:"type:varchar(255);primaryKey;column:token"`
	ExpiresAt    time.Time `gorm:"column:expires_at"`
	IsRemembered bool      `gorm:"column:is_remembered"`
	LastActivity time.Time `gorm:"column:last_activity"`
}

// BeforeCreate is a GORM hook to set the token before creating a record
func (s *UserSession) BeforeCreate(tx *gorm.DB) (err error) {
	if s.Token == "" {
		s.Token = uuid.NewString()
	}
	return
}

// ToDomainModel converts the repository model to a domain model
func (s *UserSession) ToDomainModel() *models.UserSession {
	return &models.UserSession{
		UserID:       s.UserID,
		Token:        s.Token,
		ExpiresAt:    s.ExpiresAt,
		IsRemembered: s.IsRemembered,
		LastActivity: s.LastActivity,
	}
}

// FromDomainModel creates a repository model from a domain model
func FromUserSessionDomainModel(session *models.UserSession) *UserSession {
	return &UserSession{
		UserID:       session.UserID,
		Token:        session.Token,
		ExpiresAt:    session.ExpiresAt,
		IsRemembered: session.IsRemembered,
		LastActivity: session.LastActivity,
	}
}
