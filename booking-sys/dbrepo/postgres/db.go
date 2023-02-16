// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package dbrepo

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
	if q.createRoleStmt, err = db.PrepareContext(ctx, CreateRole); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRole: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, CreateUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, DeleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getRoleStmt, err = db.PrepareContext(ctx, GetRole); err != nil {
		return nil, fmt.Errorf("error preparing query GetRole: %w", err)
	}
	if q.getRolesStmt, err = db.PrepareContext(ctx, GetRoles); err != nil {
		return nil, fmt.Errorf("error preparing query GetRoles: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, GetUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByPhoneStmt, err = db.PrepareContext(ctx, GetUserByPhone); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByPhone: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, ListUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.updateRoleStmt, err = db.PrepareContext(ctx, UpdateRole); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateRole: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, UpdateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createRoleStmt != nil {
		if cerr := q.createRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRoleStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getRoleStmt != nil {
		if cerr := q.getRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRoleStmt: %w", cerr)
		}
	}
	if q.getRolesStmt != nil {
		if cerr := q.getRolesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRolesStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserByPhoneStmt != nil {
		if cerr := q.getUserByPhoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByPhoneStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.updateRoleStmt != nil {
		if cerr := q.updateRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateRoleStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
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
	db                 DBTX
	tx                 *sql.Tx
	createRoleStmt     *sql.Stmt
	createUserStmt     *sql.Stmt
	deleteUserStmt     *sql.Stmt
	getRoleStmt        *sql.Stmt
	getRolesStmt       *sql.Stmt
	getUserStmt        *sql.Stmt
	getUserByPhoneStmt *sql.Stmt
	listUsersStmt      *sql.Stmt
	updateRoleStmt     *sql.Stmt
	updateUserStmt     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                 tx,
		tx:                 tx,
		createRoleStmt:     q.createRoleStmt,
		createUserStmt:     q.createUserStmt,
		deleteUserStmt:     q.deleteUserStmt,
		getRoleStmt:        q.getRoleStmt,
		getRolesStmt:       q.getRolesStmt,
		getUserStmt:        q.getUserStmt,
		getUserByPhoneStmt: q.getUserByPhoneStmt,
		listUsersStmt:      q.listUsersStmt,
		updateRoleStmt:     q.updateRoleStmt,
		updateUserStmt:     q.updateUserStmt,
	}
}
