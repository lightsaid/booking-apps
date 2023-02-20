package main

import (
	"toolkit/dberr"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type createSeatRequest struct {
	HallID    int64  `json:"hall_id" binding:"required,min=1"`
	ColNumber int32  `json:"col_number" binding:"required,min=1"`
	RowNumber int32  `json:"row_number" binding:"required,min=1"`
	Status    string `json:"status" binding:"required,len=1"`
}

func (s *Server) createSeat(c *gin.Context) {
	var req createSeatRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	t, err := s.store.CreateSeat(c.Request.Context(), dbrepo.CreateSeatParams{
		HallID:    req.HallID,
		ColNumber: req.ColNumber,
		RowNumber: req.RowNumber,
		Status:    req.Status,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, t)
}

func (s *Server) listSeats(c *gin.Context) {
	var req pagingRequrest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	list, err := s.store.ListSeats(c, dbrepo.ListSeatsParams{Limit: req.PageSize, Offset: req.GetPageNum()})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, list)
}

func (s *Server) getSeat(c *gin.Context) {
	var req idUriRequest
	if ok := app.BindRequestUri(c, &req); !ok {
		return
	}
	t, err := s.store.GetSeat(c.Request.Context(), req.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

type updateSeatRequest struct {
	HallID    int64  `json:"hall_id" binding:"required,min=1"`
	ColNumber int32  `json:"col_number" binding:"required,min=1"`
	RowNumber int32  `json:"row_number" binding:"required,min=1"`
	Status    string `json:"status" binding:"required,len=1"`
}

func (s *Server) updateSeat(c *gin.Context) {
	var req updateSeatRequest
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

	t, err := s.store.UpdateSeat(c.Request.Context(), dbrepo.UpdateSeatParams{
		ID:        uri.ID,
		HallID:    req.HallID,
		ColNumber: req.ColNumber,
		RowNumber: req.RowNumber,
		Status:    req.Status,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

func (s *Server) delSeat(c *gin.Context) {
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	_, err := s.store.DeleteSeat(c.Request.Context(), uri.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, nil)
}
