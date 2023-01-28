-- name: CreateFunction :one
INSERT INTO functions(creator_id, name, language, timeout, file_ref)
VALUES (?, ?, ?, ?, ?)
RETURNING *;