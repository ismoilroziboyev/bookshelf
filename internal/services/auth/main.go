package auth

import "github.com/ismoilroziboyev/bookshelf/internal/repository"

type service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *service {
	return &service{
		repo: repo,
	}
}
