package repositories

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/eijiosakabe/ai-todos/api/internal/domain/models"
	"github.com/eijiosakabe/ai-todos/api/internal/domain/repositories"
	"github.com/eijiosakabe/ai-todos/api/internal/infrastructure/repomodels"
)

type PostgresUserSessionRepository struct {
	db *gorm.DB
}

func NewPostgresUserSessionRepository(db *gorm.DB) repositories.UserSessionRepository {
	return &PostgresUserSessionRepository{db: db}
}

func (r *PostgresUserSessionRepository) CreateSession(ctx context.Context, session *models.UserSession) error {
	repoSession := repomodels.FromUserSessionDomainModel(session)
	result := r.db.WithContext(ctx).Create(repoSession)
	return result.Error
}

func (r *PostgresUserSessionRepository) FindSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	var repoSession repomodels.UserSession
	result := r.db.WithContext(ctx).
		Where("token = ? AND expires_at > NOW()", token).
		First(&repoSession)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding session: %w", result.Error)
	}

	return repoSession.ToDomainModel(), nil
}

func (r *PostgresUserSessionRepository) DeleteSession(ctx context.Context, token string) error {
	result := r.db.WithContext(ctx).
		Where("token = ?", token).
		Delete(&repomodels.UserSession{})
	return result.Error
}

func (r *PostgresUserSessionRepository) DeleteExpiredSessions(ctx context.Context) error {
	result := r.db.WithContext(ctx).
		Where("expires_at <= NOW()").
		Delete(&repomodels.UserSession{})
	return result.Error
}
