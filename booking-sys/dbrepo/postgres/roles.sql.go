// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: roles.sql

package dbrepo

import (
	"context"
)

const CreateRole = `-- name: CreateRole :one
INSERT INTO tb_roles (name, code, description)
VALUES ($1, $2, $3) RETURNING id, name, code, description, created_at, updated_at, deleted_at
`

type CreateRoleParams struct {
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Description *string `json:"description"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (*TbRole, error) {
	row := q.queryRow(ctx, q.createRoleStmt, CreateRole, arg.Name, arg.Code, arg.Description)
	var i TbRole
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const GetRole = `-- name: GetRole :one
SELECT id, name, code, description, created_at, updated_at, deleted_at FROM tb_roles WHERE id = $1 AND deleted_at IS NULL LIMIT 1
`

func (q *Queries) GetRole(ctx context.Context, id int64) (*TbRole, error) {
	row := q.queryRow(ctx, q.getRoleStmt, GetRole, id)
	var i TbRole
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const GetRoles = `-- name: GetRoles :many
SELECT id, name, code, description, created_at, updated_at, deleted_at FROM tb_roles WHERE deleted_at IS NULL ORDER BY id LIMIT $1 OFFSET $2
`

type GetRolesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetRoles(ctx context.Context, arg GetRolesParams) ([]*TbRole, error) {
	rows, err := q.query(ctx, q.getRolesStmt, GetRoles, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*TbRole{}
	for rows.Next() {
		var i TbRole
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Description,
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

const UpdateRole = `-- name: UpdateRole :one
UPDATE tb_roles 
SET
    name = $2,
    description = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, name, code, description, created_at, updated_at, deleted_at
`

type UpdateRoleParams struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) (*TbRole, error) {
	row := q.queryRow(ctx, q.updateRoleStmt, UpdateRole, arg.ID, arg.Name, arg.Description)
	var i TbRole
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
