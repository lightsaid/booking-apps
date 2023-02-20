package main

import (
	"toolkit/dberr"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type createTheaterRequest struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

func (s *Server) createTheater(c *gin.Context) {
	var req createTheaterRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	t, err := s.store.CreateTheater(c.Request.Context(), dbrepo.CreateTheaterParams{
		Name:     req.Name,
		Location: &req.Location,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, t)
}

func (s *Server) listTheaters(c *gin.Context) {
	var req pagingRequrest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	list, err := s.store.ListTheaters(c, dbrepo.ListTheatersParams{Limit: req.PageSize, Offset: req.GetPageNum()})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, list)
}

func (s *Server) getTheater(c *gin.Context) {
	var req idUriRequest
	if ok := app.BindRequestUri(c, &req); !ok {
		return
	}
	t, err := s.store.GetTheater(c.Request.Context(), req.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

type updateTheaterRequest struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

func (s *Server) updateTheater(c *gin.Context) {
	var req updateTheaterRequest
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

	t, err := s.store.UpdateTheater(c.Request.Context(), dbrepo.UpdateTheaterParams{
		ID:       uri.ID,
		Name:     req.Name,
		Location: &req.Location,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

func (s *Server) delTheater(c *gin.Context) {
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	_, err := s.store.DeleteTheater(c.Request.Context(), uri.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, nil)
}
