package services

import (
	"github.com/ismoilroziboyev/bookshelf/internal/config"
	"github.com/ismoilroziboyev/bookshelf/internal/ports"
	"github.com/ismoilroziboyev/bookshelf/internal/repository"
	"github.com/ismoilroziboyev/bookshelf/internal/services/auth"
	"github.com/ismoilroziboyev/bookshelf/internal/services/books"
)

type Service struct {
	ports.AuthService
	ports.BooksService
}

func New(cfg *config.Config, repo *repository.Repository) *Service {
	return &Service{
		AuthService:  auth.New(repo),
		BooksService: books.New(repo, cfg),
	}
}
