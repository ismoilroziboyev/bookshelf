package auth

import (
	"context"

	"github.com/ismoilroziboyev/bookshelf/internal/domain"
	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/ismoilroziboyev/go-pkg/mapper"
	"github.com/jackc/pgx/v4"
)

func (s *service) GetUserByKey(ctx context.Context, key string) (*domain.User, error) {
	var (
		resp domain.User
	)

	user, err := s.repo.DB.GetUserByKey(ctx, key)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.NewNotFoundErrorf("user not found")
		}
		return nil, errors.NewInternalServerErrorw("cannot get user details", err)
	}

	if err := mapper.Map(ctx, user, &resp); err != nil {
		return nil, errors.NewInternalServerErrorw("cannot map user details", err)
	}

	return &resp, nil
}
