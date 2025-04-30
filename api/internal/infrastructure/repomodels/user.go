package repomodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/eijiosakabe/ai-todos/api/internal/domain/models"
)

type User struct {
	gorm.Model
	ID           string    `gorm:"type:uuid;primaryKey;column:id"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null;column:email"`
	PasswordHash string    `gorm:"type:varchar(255);not null;column:password_hash"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

// BeforeCreate is a GORM hook to set the UUID before creating a record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}

// ToDomainModel converts the repository model to a domain model
func (u *User) ToDomainModel() *models.User {
	return &models.User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

// FromDomainModel creates a repository model from a domain model
func FromDomainModel(user *models.User) *User {
	return &User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
