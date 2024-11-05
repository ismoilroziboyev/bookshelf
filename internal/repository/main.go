package repository

import (
	"context"

	"github.com/ismoilroziboyev/bookshelf/internal/config"
	"github.com/ismoilroziboyev/bookshelf/internal/repository/postgres"
)

type Repository struct {
	DB postgres.Database
}

func New(cfg *config.Config) *Repository {
	return &Repository{
		DB: postgres.New(cfg),
	}
}

func (r *Repository) Close(ctx context.Context) error {
	r.DB.Close()
	return nil
}
