-- name: CreateShowtime :one
INSERT INTO tb_showtimes(
    movie_id,
    hall_id,
    start_time,
    end_time
)
VALUES($1, $2, $3, $4) RETURNING *;

-- name: GetShowtime :one
SELECT * FROM tb_showtimes WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: ListShowtimes :many
SELECT * FROM tb_showtimes WHERE deleted_at IS NULL ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: DeleteShowtime :one
UPDATE tb_showtimes SET deleted_at = now() WHERE id = $1 AND deleted_at IS NOT NULL RETURNING *;;

-- name: UpdateShowtime :one 
UPDATE tb_showtimes SET 
    movie_id = $2,
    hall_id = $3,
    start_time = $4,
    end_time = $5
WHERE id = $1 AND deleted_at IS NOT NULL
RETURNING *;