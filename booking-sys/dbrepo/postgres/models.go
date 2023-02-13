// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package dbrepo

import (
	"database/sql"
	"time"
)

type TbHall struct {
	ID         int64         `db:"id" json:"id"`
	TheaterID  int64         `db:"theater_id" json:"theater_id"`
	Name       string        `db:"name" json:"name"`
	TotalSeats sql.NullInt32 `db:"total_seats" json:"total_seats"`
	CreatedAt  time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time     `db:"updated_at" json:"updated_at"`
	DeletedAt  sql.NullTime  `db:"deleted_at" json:"deleted_at"`
}

type TbMovie struct {
	ID          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	ReleaseDate time.Time `db:"release_date" json:"release_date"`
	// 导演
	Director string `db:"director" json:"director"`
	// 海报/封面
	Poster string `db:"poster" json:"poster"`
	// 时长，单位: 分钟
	Duration int32 `db:"duration" json:"duration"`
	// 类型
	Genre sql.NullString `db:"genre" json:"genre"`
	// 主演
	Star        sql.NullString `db:"star" json:"star"`
	Description sql.NullString `db:"description" json:"description"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}

type TbPayment struct {
	ID       int64         `db:"id" json:"id"`
	UserID   sql.NullInt64 `db:"user_id" json:"user_id"`
	TicketID sql.NullInt64 `db:"ticket_id" json:"ticket_id"`
	// 电影票数
	NumberOfSeats int32     `db:"NumberOfSeats" json:"NumberOfSeats"`
	PaymentDate   time.Time `db:"payment_date" json:"payment_date"`
	// 支付方式
	PaymentMethod string `db:"payment_method" json:"payment_method"`
	// 支付总额, 单位：分
	PaymentAmount int32        `db:"payment_amount" json:"payment_amount"`
	CreatedAt     time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type TbRole struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Code        string         `db:"code" json:"code"`
	Description sql.NullString `db:"description" json:"description"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}

type TbSeat struct {
	ID        int64 `db:"id" json:"id"`
	HallID    int64 `db:"hall_id" json:"hall_id"`
	ColNumber int32 `db:"col_number" json:"col_number"`
	RowNumber int32 `db:"row_number" json:"row_number"`
	// 状态: A、B、N, 分别是: 可用、被预订、损坏
	Status    string       `db:"status" json:"status"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type TbShowtime struct {
	ID      int64 `db:"id" json:"id"`
	MovieID int64 `db:"movie_id" json:"movie_id"`
	HallID  int64 `db:"hall_id" json:"hall_id"`
	// 放映时间
	StartTime time.Time `db:"start_time" json:"start_time"`
	// 结束时间
	EndTime   time.Time    `db:"end_time" json:"end_time"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type TbTheater struct {
	ID        int64          `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	Location  sql.NullString `db:"location" json:"location"`
	CreatedAt time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}

type TbTicket struct {
	ID         int64         `db:"id" json:"id"`
	UserID     sql.NullInt64 `db:"user_id" json:"user_id"`
	ShowtimeID int64         `db:"showtime_id" json:"showtime_id"`
	SeatID     int64         `db:"seat_id" json:"seat_id"`
	// 单价，单位: 分
	Price int32 `db:"price" json:"price"`
	// 下订日期，被预订时设置时间
	BookingDate sql.NullTime `db:"booking_date" json:"booking_date"`
	// 支付状态:Y/N
	PaymentStatus sql.NullString `db:"payment_status" json:"payment_status"`
	CreatedAt     time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt     sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}

type TbUser struct {
	ID          int64          `db:"id" json:"id"`
	RoleID      int64          `db:"role_id" json:"role_id"`
	PhoneNumber string         `db:"phone_number" json:"phone_number"`
	Password    sql.NullString `db:"password" json:"password"`
	Name        string         `db:"name" json:"name"`
	Avatar      sql.NullString `db:"avatar" json:"avatar"`
	Openid      sql.NullString `db:"openid" json:"openid"`
	Unionid     sql.NullString `db:"unionid" json:"unionid"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}