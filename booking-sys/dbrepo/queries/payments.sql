-- name: CreatePayment :one 
INSERT INTO tb_payments(
    "user_id", "ticket_id", 
    "NumberOfSeats", "payment_date", 
    "payment_method", "payment_amount"
) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetPayment :one
SELECT * FROM tb_payments WHERE id=$1 AND deleted_at IS NULL LIMIT 1;

-- name: ListPayments :many
SELECT * FROM tb_payments WHERE deleted_at IS NULL LIMIT $1 OFFSET $2;

-- name: UpdatePayment :one
UPDATE tb_payments
SET 
    "user_id"=$2, "ticket_id"=$3, 
    "NumberOfSeats"=$4, "payment_date"=$5,
    "payment_method"=$6, "payment_amount"=$7
WHERE id=$1 AND  deleted_at IS NULL RETURNING *;

-- name: DeletePayment :one
UPDATE tb_payments
SET deleted_at = now() WHERE  id = $1 AND  deleted_at IS NULL RETURNING *;