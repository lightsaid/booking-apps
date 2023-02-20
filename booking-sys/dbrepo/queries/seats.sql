-- name: CreateSeat :one 
INSERT INTO tb_seats(
    "hall_id", "col_number", "row_number", "status"
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetSeat :one
SELECT * FROM tb_seats WHERE id=$1 AND deleted_at IS NULL LIMIT 1;

-- name: ListSeats :many
SELECT * FROM tb_seats WHERE deleted_at IS NULL LIMIT $1 OFFSET $2;

-- name: UpdateSeat :one
UPDATE tb_seats
SET 
    "hall_id"=$2, "col_number"=$3, "row_number"=$4, "status"=$5
WHERE id=$1 AND  deleted_at IS NULL RETURNING *;

-- name: DeleteSeat :one
UPDATE tb_seats
SET deleted_at = now() WHERE id = $1 AND deleted_at IS NULL RETURNING *;