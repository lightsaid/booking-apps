package main

import (
	"time"
	"toolkit/dberr"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type createPaymentRequest struct {
	UserID        *int64    `json:"user_id" binding:"-"`
	TicketID      *int64    `json:"ticket_id" binding:"required,min=1"`
	NumberOfSeats int32     `json:"NumberOfSeats" binding:"required,min=0"`
	PaymentDate   time.Time `json:"payment_date" binding:"required"`
	PaymentMethod string    `json:"payment_method" binding:"required"`
	PaymentAmount int32     `json:"payment_amount" binding:"required,min=0"`
}

func (s *Server) createPayment(c *gin.Context) {
	var req createPaymentRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	t, err := s.store.CreatePayment(c.Request.Context(), dbrepo.CreatePaymentParams{
		UserID:        req.UserID,
		TicketID:      req.TicketID,
		NumberOfSeats: req.NumberOfSeats,
		PaymentDate:   req.PaymentDate,
		PaymentMethod: req.PaymentMethod,
		PaymentAmount: req.PaymentAmount,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, t)
}

func (s *Server) listPayments(c *gin.Context) {
	var req pagingRequrest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	list, err := s.store.ListPayments(c, dbrepo.ListPaymentsParams{Limit: req.PageSize, Offset: req.GetPageNum()})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, list)
}

func (s *Server) getPayment(c *gin.Context) {
	var req idUriRequest
	if ok := app.BindRequestUri(c, &req); !ok {
		return
	}
	t, err := s.store.GetPayment(c.Request.Context(), req.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

func (s *Server) updatePayment(c *gin.Context) {
	var req createPaymentRequest
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

	t, err := s.store.UpdatePayment(c.Request.Context(), dbrepo.UpdatePaymentParams{
		ID:            uri.ID,
		UserID:        req.UserID,
		TicketID:      req.TicketID,
		NumberOfSeats: req.NumberOfSeats,
		PaymentDate:   req.PaymentDate,
		PaymentMethod: req.PaymentMethod,
		PaymentAmount: req.PaymentAmount,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

func (s *Server) delPayment(c *gin.Context) {
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	_, err := s.store.DeletePayment(c.Request.Context(), uri.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, nil)
}
