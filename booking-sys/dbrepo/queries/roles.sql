-- name: GetRole :one
SELECT * FROM tb_roles WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetRoles :many
SELECT * FROM tb_roles WHERE deleted_at IS NULL ORDER BY id LIMIT $1 OFFSET $2;
