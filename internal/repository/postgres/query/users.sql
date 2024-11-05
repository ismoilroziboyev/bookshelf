-- name: GetUserByKey :one
SELECT * FROM users WHERE key=$1;

-- name: CreateNewUser :one
INSERT INTO users (
    name,
    email,
    key,
    secret
) VALUES (
    $1,$2,$3,$4
) RETURNING *;