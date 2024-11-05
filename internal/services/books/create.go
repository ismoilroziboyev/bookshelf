package books

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ismoilroziboyev/bookshelf/internal/domain"
	"github.com/ismoilroziboyev/bookshelf/internal/repository/postgres/sqlc"
	"github.com/ismoilroziboyev/go-pkg/errors"
)

func (s *service) Create(ctx context.Context, userID int64, isbn string) (*domain.R, error) {

	args, err := s.getBookInfoISBN(isbn)

	if err != nil {
		return nil, errors.Wrap("cannot get book details from openlibrary", err)
	}

	args.UserID = userID

	book, err := s.repo.DB.CreateBook(ctx, *args)

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

func (s *service) getBookInfoISBN(isbn string) (*sqlc.CreateBookParams, error) {
	var result struct {
		Title       string                   `json:"title"`
		Authors     []map[string]interface{} `json:"authors"`
		Pages       int32                    `json:"number_of_pages"`
		Covers      []int64                  `json:"covers"`
		PublishDate string                   `json:"publish_date"`
	}

	res, err := s.resty.R().SetResult(&result).Get(fmt.Sprintf("/isbn/%s.json", isbn))

	if err != nil {
		return nil, errors.NewInternalServerErrorw("cannot make request to openlibrary", err)
	}

	if res.StatusCode() != http.StatusOK {
		if res.StatusCode() == http.StatusNotFound {
			return nil, errors.NewNotFoundErrorf("book not found with isbn")
		}
		return nil, errors.NewInternalServerErrorf("cannot make sucessfull request to openlibrary: %d", res.StatusCode())
	}

	p, err := strconv.Atoi(result.PublishDate[len(result.PublishDate)-4:])

	if err != nil {
		return nil, errors.NewInternalServerErrorw("cannot parse published date", err)
	}

	author := ""

	if len(result.Authors) != 0 {
		a, err := s.getAuthorInfo(result.Authors[0]["key"].(string))

		if err != nil {
			return nil, errors.Wrap("cannot get author details", err)
		}

		author = a
	}

	cover := ""

	if len(result.Covers) != 0 {
		cover = fmt.Sprintf("https://covers.openlibrary.org/b/id/%d-M.jpg", result.Covers[0])
	}

	return &sqlc.CreateBookParams{
		Title:     result.Title,
		Status:    0,
		Pages:     result.Pages,
		Author:    author,
		Published: int32(p),
		Isbn:      isbn,
		Cover:     cover,
	}, nil
}

func (s *service) getAuthorInfo(path string) (string, error) {

	var result struct {
		Name string `json:"name"`
	}

	res, err := s.resty.R().SetResult(&result).Get(fmt.Sprintf("%s.json", path))

	if err != nil {
		return "", errors.NewInternalServerErrorw("cannot make request to openlibrary", err)
	}

	if res.StatusCode() != http.StatusOK {
		if res.StatusCode() == http.StatusNotFound {
			return "", errors.NewNotFoundErrorf("author not found with isbn")
		}
		return "", errors.NewInternalServerErrorf("cannot make sucessfull request to openlibrary: %d", res.StatusCode())
	}

	return result.Name, nil

}
