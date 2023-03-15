-- name: CreateUser :one
INSERT INTO tb_users (
    phone_number,
    name
) VALUES($1, $2) RETURNING *;

-- name: GetUser :one
SELECT * FROM tb_users WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: UpdateUser :one
UPDATE tb_users 
SET 
    name = $2,
    avatar = $3,
    openid = $4,
    unionid = $5,
    updated_at = $6
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: GetUserByPhone :one
SELECT * FROM tb_users WHERE phone_number = $1 AND deleted_at IS NULL LIMIT 1;