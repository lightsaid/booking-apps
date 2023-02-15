package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	// if err := c.ShouldBind(&req); err != nil {
	// 	e := errs.InvalidParams.AsException(err, err.Error())
	// 	app.ToErrorResponse(c, e)
	// 	return
	// }

	if ok := app.BindRequest(c, &req); !ok {
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
