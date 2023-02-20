package main

import (
	"net/http"
	"strconv"
	"toolkit/dberr"
	"toolkit/errs"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type createUserRequest struct {
	PhoneNumber string `json:"phone_number" zh:"手机号码" binding:"required,len=11,vphone"`
	Password    string `json:"password" zh:"密码" binding:"required"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Openid      string `json:"openid"`
	Unionid     string `json:"unionid"`
}

func (s *Server) createUser(c *gin.Context) {
	// s.store.CreateUser()
	var req createUserRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	c.String(http.StatusOK, "Create User.")

}

func (s *Server) updateUser(c *gin.Context) {
	c.String(http.StatusOK, "Update User.")
}

func (s *Server) getListUsers(c *gin.Context) {
	var req pagingRequrest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

	list, err := s.store.ListUsers(c, dbrepo.ListUsersParams{Limit: req.PageSize, Offset: (req.PageNum - 1) * req.PageSize})
	if err != nil {
		app.ToErrorResponse(c, errs.ServerError.AsException(err))
		return
	}

	app.ToResponse(c, list)
}

func (s *Server) getUserById(c *gin.Context) {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		app.ToErrorResponse(c, errs.InvalidParams.AsException(err))
		return
	}

	user, err := s.store.GetUser(c, int64(id))
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	app.ToResponse(c, user)
}
