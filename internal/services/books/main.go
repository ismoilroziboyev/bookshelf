package books

import (
	"github.com/go-resty/resty/v2"
	"github.com/ismoilroziboyev/bookshelf/internal/config"
	"github.com/ismoilroziboyev/bookshelf/internal/repository"
)

type service struct {
	repo  *repository.Repository
	resty *resty.Client
}

func New(repo *repository.Repository, cfg *config.Config) *service {
	return &service{
		repo:  repo,
		resty: resty.New().SetDebug(true).SetBaseURL("https://openlibrary.org"),
	}
}
