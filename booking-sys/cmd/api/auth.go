package main

import (
	"fmt"
	"net/http"
	"time"
	"toolkit/errs"
	"toolkit/jwtutil"
	"toolkit/mocksms"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type userResponse struct {
	ID          int64     `json:"id"`
	RoleID      int64     `json:"role_id"`
	PhoneNumber string    `json:"phone_number"`
	Name        string    `son:"name"`
	Avatar      string    `json:"avatar"`
	Openid      string    `json:"-"`
	Unionid     string    `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"-"`
}

func (rsp *userResponse) toUserResponse(user *dbrepo.TbUser) *userResponse {
	rsp.ID = user.ID
	rsp.RoleID = user.RoleID
	rsp.PhoneNumber = user.PhoneNumber
	rsp.Name = user.Name
	rsp.Avatar = user.Avatar.String
	rsp.Openid = user.Openid.String
	rsp.Unionid = user.Unionid.String
	rsp.CreatedAt = user.CreatedAt
	rsp.UpdatedAt = user.UpdatedAt
	rsp.DeletedAt = user.DeletedAt.Time
	return rsp
}

type loginUserRequest struct {
	PhoneNumber string `json:"phone_number" zh:"手机号码" binding:"required,len=11"`
	Code        int64  `json:"code" zh:"验证码" binding:"required,min=1000,max=9999"`
}

type loginUserResponse struct {
	userResponse
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *Server) loginUser(c *gin.Context) {
	var req loginUserRequest
	if ok := app.BindRequest(c, &req); !ok {
		return
	}
	//查找用户是否存在
	user, err := s.store.GetUserByPhone(c.Request.Context(), req.PhoneNumber)
	if err != nil {
		app.ToErrorResponse(c, errs.NotFound)
		return
	}

	if s.config.Server.RunMode == "release" {
		// TODO: 真实发短信
		fmt.Println("待实现接入第三方短信接口。")
	} else {
		// 验证短信验证码
		ss, ok := mocksms.GetMockSMS(req.PhoneNumber)
		if !ok || ss.Code() != req.Code {
			app.ToErrorResponse(c, errs.InvalidParams.AsMessage("验证码不匹配"))
			return
		}
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "booking sys",
		Subject:   "Booking Apps",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.JWT.AccessTokenDuration)),
	}

	// 生成 access token
	accessToken, err := s.jwt.GenToken(jwtutil.NewJWTPayload(user.ID, claims))
	if err != nil {
		app.ToErrorResponse(c, errs.ServerError.AsException(err))
		return
	}

	// 生成 refresh token
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(s.config.RefreshTokenDuration))
	refreshToken, err := s.jwt.GenToken(jwtutil.NewJWTPayload(user.ID, claims))
	if err != nil {
		app.ToErrorResponse(c, errs.ServerError.AsException(err))
		return
	}

	// 响应
	var ursp userResponse
	ursp.toUserResponse(user)
	response := loginUserResponse{
		userResponse: *ursp.toUserResponse(user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	app.ToResponse(c, response)
}

func (s *Server) refreshToken(c *gin.Context) {
	c.String(http.StatusOK, "Refresh Token.")
}
