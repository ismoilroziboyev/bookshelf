package books

import (
	"context"

	"github.com/ismoilroziboyev/bookshelf/internal/repository/postgres/sqlc"
	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/jackc/pgx/v4"
)

func (s *service) Delete(ctx context.Context, userID, bookID int64) error {

	if _, err := s.repo.DB.DeleteBook(ctx, sqlc.DeleteBookParams{
		ID:     bookID,
		UserID: userID,
	}); err != nil {
		if err == pgx.ErrNoRows {
			return errors.NewNotFoundErrorf("book not found")
		}
		return errors.NewInternalServerErrorw("cannot delete book", err)
	}

	return nil
}
