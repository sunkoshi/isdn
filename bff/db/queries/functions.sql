-- name: CreateFunction :one
INSERT INTO functions(creator_id, name, language, timeout, output_type)
VALUES (?, ?, ?, ?, ?)
RETURNING *;
-- name: GetFunctionsByCreatorId :many
SELECT *
FROM functions
WHERE creator_id = ?;
-- name: GetFunctionsById :one
SELECT *
FROM functions
WHERE id = ?;
-- name: DeleteFunctionsByIdAndCreatorId :exec
DELETE FROM functions
WHERE id = ?
    AND creator_id = ?;