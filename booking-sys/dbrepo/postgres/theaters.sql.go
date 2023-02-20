// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: theaters.sql

package dbrepo

import (
	"context"
)

const CreateTheater = `-- name: CreateTheater :one
INSERT INTO tb_theaters("name", "location") VALUES($1, $2) RETURNING id, name, location, created_at, updated_at, deleted_at
`

type CreateTheaterParams struct {
	Name     string  `json:"name"`
	Location *string `json:"location"`
}

func (q *Queries) CreateTheater(ctx context.Context, arg CreateTheaterParams) (*TbTheater, error) {
	row := q.queryRow(ctx, q.createTheaterStmt, CreateTheater, arg.Name, arg.Location)
	var i TbTheater
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const DeleteTheater = `-- name: DeleteTheater :one
UPDATE tb_theaters
SET deleted_at = now() WHERE  id = $1 AND deleted_at IS NULL RETURNING id, name, location, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteTheater(ctx context.Context, id int64) (*TbTheater, error) {
	row := q.queryRow(ctx, q.deleteTheaterStmt, DeleteTheater, id)
	var i TbTheater
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const GetTheater = `-- name: GetTheater :one
SELECT id, name, location, created_at, updated_at, deleted_at FROM tb_theaters WHERE id = $1 AND deleted_at IS NULL LIMIT 1
`

func (q *Queries) GetTheater(ctx context.Context, id int64) (*TbTheater, error) {
	row := q.queryRow(ctx, q.getTheaterStmt, GetTheater, id)
	var i TbTheater
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const ListTheaters = `-- name: ListTheaters :many
SELECT id, name, location, created_at, updated_at, deleted_at FROM tb_theaters WHERE deleted_at IS NULL LIMIT $1 OFFSET $2
`

type ListTheatersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTheaters(ctx context.Context, arg ListTheatersParams) ([]*TbTheater, error) {
	rows, err := q.query(ctx, q.listTheatersStmt, ListTheaters, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*TbTheater{}
	for rows.Next() {
		var i TbTheater
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Location,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateTheater = `-- name: UpdateTheater :one
UPDATE tb_theaters 
SET "name" = $2, "location" = $3 WHERE id = $1 AND  deleted_at IS NULL RETURNING id, name, location, created_at, updated_at, deleted_at
`

type UpdateTheaterParams struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Location *string `json:"location"`
}

func (q *Queries) UpdateTheater(ctx context.Context, arg UpdateTheaterParams) (*TbTheater, error) {
	row := q.queryRow(ctx, q.updateTheaterStmt, UpdateTheater, arg.ID, arg.Name, arg.Location)
	var i TbTheater
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
