package auth

import (
	"context"

	"github.com/ismoilroziboyev/bookshelf/internal/domain"
	"github.com/ismoilroziboyev/bookshelf/internal/repository/postgres/sqlc"
	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/jackc/pgx/v4"
)

func (s *service) SignUp(ctx context.Context, payload *domain.SignUpPayload) (*domain.R, error) {

	_, err := s.repo.DB.GetUserByKey(ctx, payload.Key)

	if err != nil && err != pgx.ErrNoRows {
		return nil, errors.NewInternalServerErrorw("cannot get user by key", err)
	}

	if err == nil {
		return nil, errors.NewBadRequestErrorf("user is exists with this key")
	}

	newUser, err := s.repo.DB.CreateNewUser(ctx, sqlc.CreateNewUserParams{
		Name:   payload.Name,
		Email:  payload.Email,
		Key:    payload.Key,
		Secret: payload.Secret,
	})

	if err != nil {
		return nil, errors.NewInternalServerErrorw("cannot create new user", err)
	}

	return &domain.R{
		"id":    newUser.ID,
		"name":  newUser.Name,
		"email": newUser.Email,
		"key":   newUser.Key,
	}, nil
}
