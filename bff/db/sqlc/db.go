// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createFunctionStmt, err = db.PrepareContext(ctx, createFunction); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFunction: %w", err)
	}
	if q.createFunctionCallStmt, err = db.PrepareContext(ctx, createFunctionCall); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFunctionCall: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteFunctionsByIdAndCreatorIdStmt, err = db.PrepareContext(ctx, deleteFunctionsByIdAndCreatorId); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteFunctionsByIdAndCreatorId: %w", err)
	}
	if q.getFunctionByIdStmt, err = db.PrepareContext(ctx, getFunctionById); err != nil {
		return nil, fmt.Errorf("error preparing query GetFunctionById: %w", err)
	}
	if q.getFunctionsByCreatorIdStmt, err = db.PrepareContext(ctx, getFunctionsByCreatorId); err != nil {
		return nil, fmt.Errorf("error preparing query GetFunctionsByCreatorId: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.updateFunctionCallStmt, err = db.PrepareContext(ctx, updateFunctionCall); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateFunctionCall: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createFunctionStmt != nil {
		if cerr := q.createFunctionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFunctionStmt: %w", cerr)
		}
	}
	if q.createFunctionCallStmt != nil {
		if cerr := q.createFunctionCallStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFunctionCallStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteFunctionsByIdAndCreatorIdStmt != nil {
		if cerr := q.deleteFunctionsByIdAndCreatorIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteFunctionsByIdAndCreatorIdStmt: %w", cerr)
		}
	}
	if q.getFunctionByIdStmt != nil {
		if cerr := q.getFunctionByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFunctionByIdStmt: %w", cerr)
		}
	}
	if q.getFunctionsByCreatorIdStmt != nil {
		if cerr := q.getFunctionsByCreatorIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFunctionsByCreatorIdStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.updateFunctionCallStmt != nil {
		if cerr := q.updateFunctionCallStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateFunctionCallStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                  DBTX
	tx                                  *sql.Tx
	createFunctionStmt                  *sql.Stmt
	createFunctionCallStmt              *sql.Stmt
	createUserStmt                      *sql.Stmt
	deleteFunctionsByIdAndCreatorIdStmt *sql.Stmt
	getFunctionByIdStmt                 *sql.Stmt
	getFunctionsByCreatorIdStmt         *sql.Stmt
	getUserStmt                         *sql.Stmt
	getUserByEmailStmt                  *sql.Stmt
	updateFunctionCallStmt              *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                  tx,
		tx:                                  tx,
		createFunctionStmt:                  q.createFunctionStmt,
		createFunctionCallStmt:              q.createFunctionCallStmt,
		createUserStmt:                      q.createUserStmt,
		deleteFunctionsByIdAndCreatorIdStmt: q.deleteFunctionsByIdAndCreatorIdStmt,
		getFunctionByIdStmt:                 q.getFunctionByIdStmt,
		getFunctionsByCreatorIdStmt:         q.getFunctionsByCreatorIdStmt,
		getUserStmt:                         q.getUserStmt,
		getUserByEmailStmt:                  q.getUserByEmailStmt,
		updateFunctionCallStmt:              q.updateFunctionCallStmt,
	}
}
