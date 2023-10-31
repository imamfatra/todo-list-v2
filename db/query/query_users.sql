-- name: CreateAccount :one
INSERT INTO users (
    email,
    username,
    password
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT *
FROM users
WHERE username = $1 AND password = $2;
