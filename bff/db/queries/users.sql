-- name: GetUser :one
SELECT *
FROM users;
-- name: GetUserByEmail :one
SELECT *
from users
WHERE email = ?;
-- name: CreateUser :one
INSERT INTO users(email, password, type)
VALUES (?, ?, ?)
RETURNING *;