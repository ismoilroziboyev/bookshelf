package ports

import (
	"context"

	"github.com/ismoilroziboyev/bookshelf/internal/domain"
)

type BooksService interface {
	Create(ctx context.Context, userID int64, isbn string) (*domain.R, error)
	GetAll(ctx context.Context, userID int64) ([]domain.R, error)
	Search(ctx context.Context, userID int64, search string) ([]domain.R, error)
	Edit(ctx context.Context, payload *domain.EditBookPayload) (*domain.R, error)
	Delete(ctx context.Context, userID, bookID int64) error
}
