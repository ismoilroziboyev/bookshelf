package ports

import (
	"context"

	"github.com/ismoilroziboyev/bookshelf/internal/domain"
)

type AuthService interface {
	SignUp(ctx context.Context, payload *domain.SignUpPayload) (*domain.R, error)
	GetUserByKey(ctx context.Context, key string) (*domain.User, error)
}
