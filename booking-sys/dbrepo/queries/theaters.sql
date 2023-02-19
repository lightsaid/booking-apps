-- name: CreateTheater :one 
INSERT INTO tb_theaters("name", "location") VALUES($1, $2) RETURNING *;

-- name: GetTheater :one
SELECT * FROM tb_theaters WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: ListTheaters :many
SELECT * FROM tb_theaters WHERE deleted_at IS NULL LIMIT $1 OFFSET $2;

-- name: UpdateTheater :one
UPDATE tb_theaters 
SET "name" = $2, "location" = $3 WHERE id = $1 AND  deleted_at IS NULL RETURNING *;

-- name: DeleteTheater :one
UPDATE tb_theaters
SET deleted_at = now() WHERE deleted_at IS NULL RETURNING *;