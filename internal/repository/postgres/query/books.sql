-- name: CreateBook :one
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
) ON CONFLICT(user_id,isbn) DO UPDATE SET updated_at=NOW() RETURNING *;


-- name: GetAllBooks :many
SELECT 
    * 
FROM books 
WHERE 
    user_id=$1
    AND
    (
        sqlc.arg(search)=''
        OR
        title ILIKE '%' || sqlc.arg(search)::VARCHAR || '%'
    )
ORDER BY created_at DESC;


-- name: UpdateBookStatus :one
UPDATE books SET status=$1, updated_at=NOW() WHERE id=$2 RETURNING *;

-- name: DeleteBook :one
DELETE FROM books WHERE id=$1 AND user_id=$2 RETURNING *;