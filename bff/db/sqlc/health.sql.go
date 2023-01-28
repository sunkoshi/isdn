// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: health.sql

package db

import (
	"context"
)

const getHealth = `-- name: GetHealth :one
SELECT NOW()
`

func (q *Queries) GetHealth(ctx context.Context) (interface{}, error) {
	row := q.queryRow(ctx, q.getHealthStmt, getHealth)
	var now interface{}
	err := row.Scan(&now)
	return now, err
}
