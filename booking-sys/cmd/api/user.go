package main

import (
	"database/sql"
	"net/http"
	"toolkit/errs"

	"github.com/gin-gonic/gin"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type createUserRequest struct {
	PhoneNumber string         `json:"phone_number" binding:"required,len=11,vphone"`
	Password    sql.NullString `json:"password"`
	Name        string         `json:"name"`
	Avatar      sql.NullString `json:"avatar"`
	Openid      sql.NullString `json:"openid"`
	Unionid     sql.NullString `json:"unionid"`
}

func (s *Server) createUser(c *gin.Context) {
	// s.store.CreateUser()
	var req createUserRequest
	if err := c.ShouldBind(&req); err != nil {
		e := errs.InvalidParams.AsException(err, err.Error())
		app.ToErrorResponse(c, e)
		return
	}

	c.String(http.StatusOK, "Create User.")
}

func (s *Server) updateUser(c *gin.Context) {
	c.String(http.StatusOK, "Update User.")
}

func (s *Server) getListUsers(c *gin.Context) {
	c.String(http.StatusOK, "List User.")
}

func (s *Server) getUserById(c *gin.Context) {
	c.String(http.StatusOK, "Get One User.")
}
