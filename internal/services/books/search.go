package books

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ismoilroziboyev/bookshelf/internal/domain"
	"github.com/ismoilroziboyev/go-pkg/errors"
)

func (s *service) Search(ctx context.Context, userID int64, search string) ([]domain.R, error) {
	var result struct {
		Docs []struct {
			AuthorName  []string `json:"author_name"`
			CoverI      int64    `json:"cover_i"`
			PublishYear []int    `json:"publish_year"`
			Title       string   `json:"title"`
			Isbn        []string `json:"isbn"`
		} `json:"docs"`
	}

	res, err := s.resty.R().SetResult(&result).Get(fmt.Sprintf("/search.json?q=%s&limit=12&offset=0", strings.ReplaceAll(search, " ", "+")))

	if err != nil {
		return nil, errors.NewInternalServerErrorw("cannot make request to openlibrary", err)
	}

	if res.StatusCode() != http.StatusOK {
		return nil, errors.NewInternalServerErrorf("cannot make successfull request to openlibrary: %d", res.StatusCode())
	}

	var resp = make([]domain.R, len(result.Docs))

	for index, book := range result.Docs {
		author := ""
		isbn := ""
		published := 0

		if len(book.AuthorName) != 0 {
			author = book.AuthorName[0]
		}

		if len(book.Isbn) != 0 {
			isbn = book.Isbn[0]
		}

		if len(book.PublishYear) != 0 {
			published = book.PublishYear[0]
		}

		resp[index] = domain.R{
			"isbn":      isbn,
			"title":     book.Title,
			"cover":     fmt.Sprintf("https://covers.openlibrary.org/b/id/%d-M.jpg", book.CoverI),
			"author":    author,
			"published": published,
		}
	}

	return resp, nil
}
