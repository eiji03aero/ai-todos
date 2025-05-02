package repositories

import (
	"context"

	"github.com/eijiosakabe/ai-todos/api/internal/domain/models"
)

type UserSessionRepository interface {
	CreateSession(ctx context.Context, session *models.UserSession) error
	FindSessionByToken(ctx context.Context, token string) (*models.UserSession, error)
	DeleteSession(ctx context.Context, token string) error
	DeleteExpiredSessions(ctx context.Context) error
}
