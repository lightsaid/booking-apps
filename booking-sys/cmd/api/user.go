package main

import (
	"strconv"
	"time"
	"toolkit/dberr"
	"toolkit/errs"
	"toolkit/jwtutil"
	"toolkit/pswd"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

const (
	ConsumerCode = "CONSUMER"
	AdminCode    = "ADMIN"
)

type createUserRequest struct {
	PhoneNumber string `json:"phone_number" zh:"手机号码" binding:"required,len=11,vphone"`
	Password    string `json:"password" binding:"-"`
	Name        string `json:"name" binding:"-"`
	Avatar      string `json:"avatar" binding:"-"`
	Openid      string `json:"openid" binding:"-"`
	Unionid     string `json:"unionid" binding:"-"`
	Role        bool   `json:"role"`
}

func (s *Server) createUser(c *gin.Context) {
	var req createUserRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	var password string
	if len(req.Password) > 0 {
		var err error
		password, err = pswd.GenerateHashPwd(req.Password)
		if err != nil {
			app.ToErrorResponse(c, errs.BadRequest.AsMessage("密码格式不对"))
			return
		}
	}
	res, err := s.store.CreateUser(c.Request.Context(), dbrepo.CreateUserParams{
		PhoneNumber: req.PhoneNumber,
		Password:    &password,
		Openid:      &req.Openid,
		Name:        req.Name,
		Avatar:      &req.Avatar,
		Unionid:     &req.Unionid,
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, res)
}

type updateUserRequest struct {
	ID     int64  `json:"id" binding:"required"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Role   bool   `json:"role"`
}

func (s *Server) updateUser(c *gin.Context) {
	var req updateUserRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	user, err := s.store.GetUser(c.Request.Context(), req.ID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	if len(req.Name) != 0 {
		user.Name = req.Name
	}

	if len(req.Avatar) != 0 {
		user.Avatar = &req.Avatar
	}

	user.UpdatedAt = time.Now()

	user, err = s.store.UpdateUser(c.Request.Context(), dbrepo.UpdateUserParams{
		ID:        user.ID,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Openid:    user.Openid,
		Unionid:   user.Unionid,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}

	var role *dbrepo.TbRole
	if req.Role {
		role, err = s.store.GetRoleByCode(c.Request.Context(), AdminCode)
		if err != nil {
			e, _ := dberr.HandleDBError(err)
			app.ToErrorResponse(c, e)
			return
		}
	} else {
		role, err = s.store.GetRoleByCode(c.Request.Context(), ConsumerCode)
		if err != nil {
			e, _ := dberr.HandleDBError(err)
			app.ToErrorResponse(c, e)
			return
		}
	}

	_, err = s.store.UpdateUserRole(c.Request.Context(), dbrepo.UpdateUserRoleParams{ID: user.ID, RoleID: &role.ID, UpdatedAt: time.Now()})
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, user)
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
		app.ToErrorResponse(c, errs.BadRequest.AsException(err))
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

func (s *Server) getProfile(c *gin.Context) {
	payload := c.MustGet(AuthorizationPayloadKey).(*jwtutil.JWTPayload)
	user, err := s.store.GetUser(c, payload.UID)
	if err != nil {
		e, _ := dberr.HandleDBError(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, user)
}
