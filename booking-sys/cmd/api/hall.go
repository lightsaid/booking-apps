package main

import (
	"toolkit/dberr"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type createHallRequest struct {
	TheaterID  int64  `json:"theater_id" binding:"required,min=1"`
	Name       string `json:"name" binding:"required"`
	TotalSeats *int32 `json:"total_seats" binding:"required"`
}

func (s *Server) createHall(c *gin.Context) {
	var req createHallRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	t, err := s.store.CreateHall(c.Request.Context(), dbrepo.CreateHallParams{
		TheaterID:  req.TheaterID,
		Name:       req.Name,
		TotalSeats: req.TotalSeats,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, t)
}

func (s *Server) listHalls(c *gin.Context) {
	var req pagingRequrest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	list, err := s.store.ListHalls(c, dbrepo.ListHallsParams{Limit: req.PageSize, Offset: req.GetPageNum()})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, list)
}

func (s *Server) getHall(c *gin.Context) {
	var req idUriRequest
	if ok := app.BindRequestUri(c, &req); !ok {
		return
	}
	t, err := s.store.GetHall(c.Request.Context(), req.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

type updateHallRequest struct {
	TheaterID  int64  `json:"theater_id" binding:"required,min=1"`
	Name       string `json:"name" binding:"required"`
	TotalSeats *int32 `json:"total_seats" binding:"required"`
}

func (s *Server) updateHall(c *gin.Context) {
	var req updateHallRequest
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

	t, err := s.store.UpdateHall(c.Request.Context(), dbrepo.UpdateHallParams{
		ID:         uri.ID,
		TheaterID:  req.TheaterID,
		Name:       req.Name,
		TotalSeats: req.TotalSeats,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

func (s *Server) delHall(c *gin.Context) {
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	_, err := s.store.DeleteHall(c.Request.Context(), uri.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, nil)
}
