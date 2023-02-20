package main

import (
	"time"
	"toolkit/dberr"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type createTicketRequest struct {
	UserID        *int64     `json:"user_id" binding:"-"`
	ShowtimeID    int64      `json:"showtime_id" binding:"required,min=1"`
	SeatID        int64      `json:"seat_id" binding:"required,min=1"`
	Price         int32      `json:"price" binding:"required,min=0"`
	BookingDate   *time.Time `json:"booking_date"`
	PaymentStatus *string    `json:"payment_status"`
}

func (s *Server) createTicket(c *gin.Context) {
	var req createTicketRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	t, err := s.store.CreateTicket(c.Request.Context(), dbrepo.CreateTicketParams{
		UserID:        req.UserID,
		ShowtimeID:    req.ShowtimeID,
		SeatID:        req.SeatID,
		Price:         req.Price,
		BookingDate:   req.BookingDate,
		PaymentStatus: req.PaymentStatus,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, t)
}

func (s *Server) listTickets(c *gin.Context) {
	var req pagingRequrest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	list, err := s.store.ListTickets(c, dbrepo.ListTicketsParams{Limit: req.PageSize, Offset: req.GetPageNum()})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, list)
}

func (s *Server) getTicket(c *gin.Context) {
	var req idUriRequest
	if ok := app.BindRequestUri(c, &req); !ok {
		return
	}
	t, err := s.store.GetTicket(c.Request.Context(), req.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

func (s *Server) updateTicket(c *gin.Context) {
	var req createTicketRequest
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

	t, err := s.store.UpdateTicket(c.Request.Context(), dbrepo.UpdateTicketParams{
		ID:            uri.ID,
		UserID:        req.UserID,
		ShowtimeID:    req.ShowtimeID,
		SeatID:        req.SeatID,
		Price:         req.Price,
		BookingDate:   req.BookingDate,
		PaymentStatus: req.PaymentStatus,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

func (s *Server) delTicket(c *gin.Context) {
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	_, err := s.store.DeleteTicket(c.Request.Context(), uri.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, nil)
}
