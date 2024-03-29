// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: function_call.sql

package db

import (
	"context"
)

const createFunctionCall = `-- name: CreateFunctionCall :one
INSERT INTO function_calls(function_id)
VALUES (?)
RETURNING id, function_id, output, stdout, error, cost, created_at
`

func (q *Queries) CreateFunctionCall(ctx context.Context, functionID int64) (*FunctionCall, error) {
	row := q.queryRow(ctx, q.createFunctionCallStmt, createFunctionCall, functionID)
	var i FunctionCall
	err := row.Scan(
		&i.ID,
		&i.FunctionID,
		&i.Output,
		&i.Stdout,
		&i.Error,
		&i.Cost,
		&i.CreatedAt,
	)
	return &i, err
}

const updateFunctionCall = `-- name: UpdateFunctionCall :one
UPDATE function_calls SET output = ?,stdout = ?,error = ?,cost=? WHERE id = ? RETURNING id, function_id, output, stdout, error, cost, created_at
`

type UpdateFunctionCallParams struct {
	Output string `json:"output"`
	Stdout string `json:"stdout"`
	Error  string `json:"error"`
	Cost   int64  `json:"cost"`
	ID     int64  `json:"id"`
}

func (q *Queries) UpdateFunctionCall(ctx context.Context, arg UpdateFunctionCallParams) (*FunctionCall, error) {
	row := q.queryRow(ctx, q.updateFunctionCallStmt, updateFunctionCall,
		arg.Output,
		arg.Stdout,
		arg.Error,
		arg.Cost,
		arg.ID,
	)
	var i FunctionCall
	err := row.Scan(
		&i.ID,
		&i.FunctionID,
		&i.Output,
		&i.Stdout,
		&i.Error,
		&i.Cost,
		&i.CreatedAt,
	)
	return &i, err
}
