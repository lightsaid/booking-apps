-- name: GetRole :one
SELECT * FROM tb_roles WHERE id = $1 AND deleted_at IS NULL LIMIT 1;


-- name: GetRoleByCode :one
SELECT * FROM tb_roles WHERE code = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetRoles :many
SELECT * FROM tb_roles WHERE deleted_at IS NULL ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateRole :one
INSERT INTO tb_roles (name, code, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateRole :one
UPDATE tb_roles 
SET
    name = $2,
    description = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;