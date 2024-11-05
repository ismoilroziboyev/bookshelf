// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package sqlc

import (
	"context"
)

type Querier interface {
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	CreateNewUser(ctx context.Context, arg CreateNewUserParams) (User, error)
	DeleteBook(ctx context.Context, arg DeleteBookParams) (Book, error)
	GetAllBooks(ctx context.Context, arg GetAllBooksParams) ([]Book, error)
	GetUserByKey(ctx context.Context, key string) (User, error)
	UpdateBookStatus(ctx context.Context, arg UpdateBookStatusParams) (Book, error)
}

var _ Querier = (*Queries)(nil)