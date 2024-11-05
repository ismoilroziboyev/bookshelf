package books

import (
	"context"

	"github.com/ismoilroziboyev/bookshelf/internal/domain"
	"github.com/ismoilroziboyev/bookshelf/internal/repository/postgres/sqlc"
	"github.com/ismoilroziboyev/go-pkg/errors"
)

func (s *service) GetAll(ctx context.Context, userID int64) ([]domain.R, error) {

	books, err := s.repo.DB.GetAllBooks(ctx, sqlc.GetAllBooksParams{
		UserID: userID,
		Search: "",
	})

	if err != nil {
		return nil, errors.NewInternalServerErrorw("cannot get books list", err)
	}

	var resp = make([]domain.R, len(books))

	for index, book := range books {
		resp[index] = domain.R{
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
		}
	}

	return resp, nil
}
