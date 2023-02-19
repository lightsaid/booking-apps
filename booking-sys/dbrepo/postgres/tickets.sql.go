// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: tickets.sql

package dbrepo

import (
	"context"
	"database/sql"
)

const CreateTicket = `-- name: CreateTicket :one
INSERT INTO tb_tickets(
    "user_id", "showtime_id", "seat_id", "price", "booking_date", "payment_status"
) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, showtime_id, seat_id, price, booking_date, payment_status, created_at, updated_at, deleted_at
`

type CreateTicketParams struct {
	UserID        sql.NullInt64  `db:"user_id" json:"user_id"`
	ShowtimeID    int64          `db:"showtime_id" json:"showtime_id"`
	SeatID        int64          `db:"seat_id" json:"seat_id"`
	Price         int32          `db:"price" json:"price"`
	BookingDate   sql.NullTime   `db:"booking_date" json:"booking_date"`
	PaymentStatus sql.NullString `db:"payment_status" json:"payment_status"`
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (*TbTicket, error) {
	row := q.queryRow(ctx, q.createTicketStmt, CreateTicket,
		arg.UserID,
		arg.ShowtimeID,
		arg.SeatID,
		arg.Price,
		arg.BookingDate,
		arg.PaymentStatus,
	)
	var i TbTicket
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ShowtimeID,
		&i.SeatID,
		&i.Price,
		&i.BookingDate,
		&i.PaymentStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const DeleteTicket = `-- name: DeleteTicket :one
UPDATE tb_tickets
SET deleted_at = now() WHERE deleted_at IS NULL RETURNING id, user_id, showtime_id, seat_id, price, booking_date, payment_status, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteTicket(ctx context.Context) (*TbTicket, error) {
	row := q.queryRow(ctx, q.deleteTicketStmt, DeleteTicket)
	var i TbTicket
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ShowtimeID,
		&i.SeatID,
		&i.Price,
		&i.BookingDate,
		&i.PaymentStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const GetTicket = `-- name: GetTicket :one
SELECT id, user_id, showtime_id, seat_id, price, booking_date, payment_status, created_at, updated_at, deleted_at FROM tb_tickets WHERE id=$1 AND deleted_at IS NULL LIMIT 1
`

func (q *Queries) GetTicket(ctx context.Context, id int64) (*TbTicket, error) {
	row := q.queryRow(ctx, q.getTicketStmt, GetTicket, id)
	var i TbTicket
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ShowtimeID,
		&i.SeatID,
		&i.Price,
		&i.BookingDate,
		&i.PaymentStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const ListTickets = `-- name: ListTickets :many
SELECT id, user_id, showtime_id, seat_id, price, booking_date, payment_status, created_at, updated_at, deleted_at FROM tb_tickets WHERE deleted_at IS NULL LIMIT $1 OFFSET $2
`

type ListTicketsParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) ListTickets(ctx context.Context, arg ListTicketsParams) ([]*TbTicket, error) {
	rows, err := q.query(ctx, q.listTicketsStmt, ListTickets, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*TbTicket{}
	for rows.Next() {
		var i TbTicket
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ShowtimeID,
			&i.SeatID,
			&i.Price,
			&i.BookingDate,
			&i.PaymentStatus,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateTicket = `-- name: UpdateTicket :one
UPDATE tb_tickets
SET 
    "user_id"=$2, "showtime_id"=$3, 
    "seat_id"=$4, "price"=$5,
    "booking_date"=$6, "payment_status"=$7
WHERE id=$1 AND  deleted_at IS NULL RETURNING id, user_id, showtime_id, seat_id, price, booking_date, payment_status, created_at, updated_at, deleted_at
`

type UpdateTicketParams struct {
	ID            int64          `db:"id" json:"id"`
	UserID        sql.NullInt64  `db:"user_id" json:"user_id"`
	ShowtimeID    int64          `db:"showtime_id" json:"showtime_id"`
	SeatID        int64          `db:"seat_id" json:"seat_id"`
	Price         int32          `db:"price" json:"price"`
	BookingDate   sql.NullTime   `db:"booking_date" json:"booking_date"`
	PaymentStatus sql.NullString `db:"payment_status" json:"payment_status"`
}

func (q *Queries) UpdateTicket(ctx context.Context, arg UpdateTicketParams) (*TbTicket, error) {
	row := q.queryRow(ctx, q.updateTicketStmt, UpdateTicket,
		arg.ID,
		arg.UserID,
		arg.ShowtimeID,
		arg.SeatID,
		arg.Price,
		arg.BookingDate,
		arg.PaymentStatus,
	)
	var i TbTicket
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ShowtimeID,
		&i.SeatID,
		&i.Price,
		&i.BookingDate,
		&i.PaymentStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
