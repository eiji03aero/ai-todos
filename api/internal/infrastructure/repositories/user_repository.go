package repositories

import (
	"context"
	"errors"

	"github.com/eijiosakabe/ai-todos/api/internal/domain/models"
	"github.com/eijiosakabe/ai-todos/api/internal/domain/repositories"
	repoModels "github.com/eijiosakabe/ai-todos/api/internal/infrastructure/repomodels"
	"gorm.io/gorm"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&repoModels.User{}).Where("email = ?", email).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	// Check email uniqueness first
	exists, err := r.EmailExists(ctx, user.Email)
	if err != nil {
		return err
	}
	if exists {
		return ErrEmailAlreadyExists
	}

	// Convert domain model to repository model
	repoUser := repoModels.FromDomainModel(user)

	// Create the user
	result := r.db.WithContext(ctx).Create(repoUser)
	if result.Error != nil {
		return result.Error
	}

	// Update the ID from the created repository model
	user.ID = repoUser.ID

	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var repoUser repoModels.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&repoUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	// Convert repository model to domain model
	return repoUser.ToDomainModel(), nil
}

// Migrate creates the users table if it doesn't exist
func (r *UserRepository) Migrate() error {
	return r.db.AutoMigrate(&repoModels.User{})
}
