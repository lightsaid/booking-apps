package main

import (
	"time"
	"toolkit/dberr"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type createMovieRequest struct {
	Title       string    `json:"title" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Director    string    `json:"director" binding:"required"`
	Poster      string    `json:"poster" binding:"required"`
	Duration    int32     `json:"duration" binding:"required"`
	Genre       *string   `json:"genre" binding:"-"`
	Star        *string   `json:"star" binding:"-"`
	Description *string   `json:"description" binding:"-"`
}

func (s *Server) createMovie(c *gin.Context) {
	var req createMovieRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	t, err := s.store.CreateMovie(c.Request.Context(), dbrepo.CreateMovieParams{
		Title:       req.Title,
		ReleaseDate: req.ReleaseDate,
		Director:    req.Director,
		Poster:      req.Poster,
		Duration:    req.Duration,
		Genre:       req.Genre,
		Star:        req.Star,
		Description: req.Description,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, t)
}

func (s *Server) listMovies(c *gin.Context) {
	var req pagingRequrest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	list, err := s.store.ListMovies(c, dbrepo.ListMoviesParams{Limit: req.PageSize, Offset: req.GetPageNum()})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, list)
}

func (s *Server) getMovie(c *gin.Context) {
	var req idUriRequest
	if ok := app.BindRequestUri(c, &req); !ok {
		return
	}
	t, err := s.store.GetMovie(c.Request.Context(), req.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

type updateMovieRequest struct {
	Title       string    `json:"title" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Director    string    `json:"director" binding:"required"`
	Poster      string    `json:"poster" binding:"required"`
	Duration    int32     `json:"duration" binding:"required"`
	Genre       *string   `json:"genre" binding:"-"`
	Star        *string   `json:"star" binding:"-"`
	Description *string   `json:"description" binding:"-"`
}

func (s *Server) updateMovie(c *gin.Context) {
	var req updateMovieRequest
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

	t, err := s.store.UpdateMovie(c.Request.Context(), dbrepo.UpdateMovieParams{
		ID:          uri.ID,
		Title:       req.Title,
		ReleaseDate: req.ReleaseDate,
		Director:    req.Director,
		Poster:      req.Poster,
		Duration:    req.Duration,
		Genre:       req.Genre,
		Star:        req.Star,
		Description: req.Description,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, t)
}

func (s *Server) delMovie(c *gin.Context) {
	var uri idUriRequest
	if ok := app.BindRequestUri(c, &uri); !ok {
		return
	}
	_, err := s.store.DeleteMovie(c.Request.Context(), uri.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, nil)
}
