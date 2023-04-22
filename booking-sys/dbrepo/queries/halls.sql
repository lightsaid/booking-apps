-- name: CreateHall :one 
INSERT INTO tb_halls("theater_id", "name", "total_seats") VALUES($1, $2, $3) RETURNING *;

-- name: GetHall :one
SELECT * FROM tb_halls WHERE id=$1 AND deleted_at IS NULL LIMIT 1;

-- name: ListHalls :many
SELECT * FROM tb_halls WHERE theater_id = $1 AND deleted_at IS NULL LIMIT $2 OFFSET $3;

-- name: UpdateHall :one
UPDATE tb_halls
SET 
    "theater_id"=$2, "name"=$3, "total_seats"=$4 
WHERE id=$1 AND  deleted_at IS NULL RETURNING *;

-- name: DeleteHall :one
UPDATE tb_halls
SET deleted_at = now() WHERE id = $1 AND  deleted_at IS NULL RETURNING *;