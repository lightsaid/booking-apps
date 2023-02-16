-- name: CreateUser :one
INSERT INTO tb_users (
    role_id,
    phone_number,
    password,
    name,
    avatar,
    openid,
    unionid
) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetUser :one
SELECT * FROM tb_users WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: ListUsers :many
SELECT * FROM tb_users WHERE deleted_at IS NULL ORDER BY id LIMIT $1 OFFSET $2;

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

-- name: DeleteUser :one
UPDATE tb_users 
    SET deleted_at = $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;


-- name: GetUserByPhone :one
SELECT * FROM tb_users WHERE phone_number = $1 LIMIT 1;

