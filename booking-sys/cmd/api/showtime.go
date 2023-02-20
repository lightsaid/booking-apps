package main

import (
	"time"
	"toolkit/dberr"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type createShowtimeRequest struct {
	MovieID   int64     `json:"movie_id" binding:"required,min=1"`
	HallID    int64     `json:"hall_id" binding:"required,min=1"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

func (s *Server) createShowtime(c *gin.Context) {
	var req createShowtimeRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	t, err := s.store.CreateShowtime(c.Request.Context(), dbrepo.CreateShowtimeParams{
		MovieID:   req.MovieID,
		HallID:    req.HallID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, t)
}

func (s *Server) listShowtimes(c *gin.Context) {
	var req pagingRequrest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	list, err := s.store.ListShowtimes(c, dbrepo.ListShowtimesParams{Limit: req.PageSize, Offset: req.GetPageNum()})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, list)
}

func (s *Server) getShowtime(c *gin.Context) {
	var req idUriRequest
	if ok := app.BindRequestUri(c, &req); !ok {
		return
	}
	t, err := s.store.GetShowtime(c.Request.Context(), req.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

type updateShowtimeRequest struct {
	MovieID   int64     `json:"movie_id" binding:"required,min=1"`
	HallID    int64     `json:"hall_id" binding:"required,min=1"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

func (s *Server) updateShowtime(c *gin.Context) {
	var req updateShowtimeRequest
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

	t, err := s.store.UpdateShowtime(c.Request.Context(), dbrepo.UpdateShowtimeParams{
		ID:        uri.ID,
		MovieID:   req.MovieID,
		HallID:    req.HallID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

func (s *Server) delShowtime(c *gin.Context) {
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	_, err := s.store.DeleteShowtime(c.Request.Context(), uri.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, nil)
}
