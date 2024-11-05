package books

import (
	"context"

	"github.com/ismoilroziboyev/bookshelf/internal/domain"
	"github.com/ismoilroziboyev/bookshelf/internal/repository/postgres/sqlc"
	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/jackc/pgx/v4"
)

func (s *service) Edit(ctx context.Context, payload *domain.EditBookPayload) (*domain.R, error) {

	book, err := s.repo.DB.UpdateBookStatus(ctx, sqlc.UpdateBookStatusParams{
		Status: payload.Status,
		ID:     payload.ID,
	})

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.NewNotFoundErrorf("book not found")
		}
		return nil, errors.NewInternalServerErrorw("cannot update book status", err)
	}

	return &domain.R{
		"book": domain.R{
			"id":        book.ID,
			"isbn":      book.Isbn,
			"title":     book.Title,
			"cover":     book.Cover,
			"author":    book.Author,
			"published": book.Published,
			"pages":     book.Pages,
		},
		"status": book.Status,
	}, nil
}
