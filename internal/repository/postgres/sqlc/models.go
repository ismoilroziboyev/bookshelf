// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package sqlc

import (
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	Isbn      string    `json:"isbn"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Author    string    `json:"author"`
	Published int32     `json:"published"`
	Pages     int32     `json:"pages"`
	Status    int32     `json:"status"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Key       string    `json:"key"`
	Secret    string    `json:"secret"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
