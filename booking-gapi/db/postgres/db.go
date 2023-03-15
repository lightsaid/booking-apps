// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

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
	if q.createUserStmt, err = db.PrepareContext(ctx, CreateUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteMovieStmt, err = db.PrepareContext(ctx, DeleteMovie); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMovie: %w", err)
	}
	if q.getMovieStmt, err = db.PrepareContext(ctx, GetMovie); err != nil {
		return nil, fmt.Errorf("error preparing query GetMovie: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, GetUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByPhoneStmt, err = db.PrepareContext(ctx, GetUserByPhone); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByPhone: %w", err)
	}
	if q.listMoviesStmt, err = db.PrepareContext(ctx, ListMovies); err != nil {
		return nil, fmt.Errorf("error preparing query ListMovies: %w", err)
	}
	if q.updateMovieStmt, err = db.PrepareContext(ctx, UpdateMovie); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMovie: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, UpdateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteMovieStmt != nil {
		if cerr := q.deleteMovieStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMovieStmt: %w", cerr)
		}
	}
	if q.getMovieStmt != nil {
		if cerr := q.getMovieStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMovieStmt: %w", cerr)
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
	if q.listMoviesStmt != nil {
		if cerr := q.listMoviesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listMoviesStmt: %w", cerr)
		}
	}
	if q.updateMovieStmt != nil {
		if cerr := q.updateMovieStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMovieStmt: %w", cerr)
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
	createUserStmt     *sql.Stmt
	deleteMovieStmt    *sql.Stmt
	getMovieStmt       *sql.Stmt
	getUserStmt        *sql.Stmt
	getUserByPhoneStmt *sql.Stmt
	listMoviesStmt     *sql.Stmt
	updateMovieStmt    *sql.Stmt
	updateUserStmt     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                 tx,
		tx:                 tx,
		createUserStmt:     q.createUserStmt,
		deleteMovieStmt:    q.deleteMovieStmt,
		getMovieStmt:       q.getMovieStmt,
		getUserStmt:        q.getUserStmt,
		getUserByPhoneStmt: q.getUserByPhoneStmt,
		listMoviesStmt:     q.listMoviesStmt,
		updateMovieStmt:    q.updateMovieStmt,
		updateUserStmt:     q.updateUserStmt,
	}
}