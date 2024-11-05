// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: books.sql

package sqlc

import (
	"context"
)

const createBook = `-- name: CreateBook :one
INSERT INTO books (
    title,
    author,
    pages,
    status,
    isbn,
    user_id,
    cover,
    published
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8
) ON CONFLICT(user_id,isbn) DO UPDATE SET updated_at=NOW() RETURNING id, isbn, title, cover, author, published, pages, status, user_id, created_at, updated_at
`

type CreateBookParams struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Pages     int32  `json:"pages"`
	Status    int32  `json:"status"`
	Isbn      string `json:"isbn"`
	UserID    int64  `json:"user_id"`
	Cover     string `json:"cover"`
	Published int32  `json:"published"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRow(ctx, createBook,
		arg.Title,
		arg.Author,
		arg.Pages,
		arg.Status,
		arg.Isbn,
		arg.UserID,
		arg.Cover,
		arg.Published,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Isbn,
		&i.Title,
		&i.Cover,
		&i.Author,
		&i.Published,
		&i.Pages,
		&i.Status,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :one
DELETE FROM books WHERE id=$1 AND user_id=$2 RETURNING id, isbn, title, cover, author, published, pages, status, user_id, created_at, updated_at
`

type DeleteBookParams struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

func (q *Queries) DeleteBook(ctx context.Context, arg DeleteBookParams) (Book, error) {
	row := q.db.QueryRow(ctx, deleteBook, arg.ID, arg.UserID)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Isbn,
		&i.Title,
		&i.Cover,
		&i.Author,
		&i.Published,
		&i.Pages,
		&i.Status,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllBooks = `-- name: GetAllBooks :many
SELECT 
    id, isbn, title, cover, author, published, pages, status, user_id, created_at, updated_at 
FROM books 
WHERE 
    user_id=$1
    AND
    (
        $2=''
        OR
        title ILIKE '%' || $2::VARCHAR || '%'
    )
ORDER BY created_at DESC
`

type GetAllBooksParams struct {
	UserID int64       `json:"user_id"`
	Search interface{} `json:"search"`
}

func (q *Queries) GetAllBooks(ctx context.Context, arg GetAllBooksParams) ([]Book, error) {
	rows, err := q.db.Query(ctx, getAllBooks, arg.UserID, arg.Search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Isbn,
			&i.Title,
			&i.Cover,
			&i.Author,
			&i.Published,
			&i.Pages,
			&i.Status,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBookStatus = `-- name: UpdateBookStatus :one
UPDATE books SET status=$1, updated_at=NOW() WHERE id=$2 RETURNING id, isbn, title, cover, author, published, pages, status, user_id, created_at, updated_at
`

type UpdateBookStatusParams struct {
	Status int32 `json:"status"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateBookStatus(ctx context.Context, arg UpdateBookStatusParams) (Book, error) {
	row := q.db.QueryRow(ctx, updateBookStatus, arg.Status, arg.ID)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Isbn,
		&i.Title,
		&i.Cover,
		&i.Author,
		&i.Published,
		&i.Pages,
		&i.Status,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
