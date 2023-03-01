-- name: CreateFunctionCall :one
INSERT INTO function_calls(function_id)
VALUES (?)
RETURNING *;
-- name: UpdateFunctionCall :one
UPDATE function_calls SET output = ?,stdout = ?,error = ?,cost=? WHERE id = ? RETURNING *;