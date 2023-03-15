-- name: GetMovie :one
SELECT * FROM tb_movies WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: ListMovies :many
SELECT * FROM tb_movies WHERE deleted_at IS NULL ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: DeleteMovie :one
UPDATE tb_movies SET deleted_at = now() WHERE id = $1 AND deleted_at IS NOT NULL RETURNING *;


-- name: UpdateMovie :one 
UPDATE tb_movies SET 
    title = $2,
    release_date = $3,
    director = $4,
    poster = $5,
    duration = $6,
    genre = $7,
    star = $8,
    "description" = $9
WHERE id = $1 AND deleted_at IS NOT NULL 
RETURNING *;