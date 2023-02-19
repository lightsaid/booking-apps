-- name: CreateTicket :one 
INSERT INTO tb_tickets(
    "user_id", "showtime_id", "seat_id", "price", "booking_date", "payment_status"
) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetTicket :one
SELECT * FROM tb_tickets WHERE id=$1 AND deleted_at IS NULL LIMIT 1;

-- name: ListTickets :many
SELECT * FROM tb_tickets WHERE deleted_at IS NULL LIMIT $1 OFFSET $2;

-- name: UpdateTicket :one
UPDATE tb_tickets
SET 
    "user_id"=$2, "showtime_id"=$3, 
    "seat_id"=$4, "price"=$5,
    "booking_date"=$6, "payment_status"=$7
WHERE id=$1 AND  deleted_at IS NULL RETURNING *;

-- name: DeleteTicket :one
UPDATE tb_tickets
SET deleted_at = now() WHERE deleted_at IS NULL RETURNING *;