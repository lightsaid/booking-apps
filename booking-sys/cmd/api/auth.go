package main

import (
	"fmt"
	"time"
	"toolkit/errs"
	"toolkit/jwtutil"
	"toolkit/mocksms"
	"toolkit/pswd"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

const (
	LoginTypeCode = "CODE"
	LoginTypePass = "PASS"
)

type loginUserRequest struct {
	PhoneNumber string `json:"phone_number" zh:"手机号码" binding:"required,len=11"`
	Code        int64  `json:"code" zh:"验证码" binding:"-"`
	Password    string `json:"password" zh:"密码" binding:"-"`
	LoginType   string `json:"login_type" zh:"登录类型" binding:"required,oneof=CODE PASS"`
}

type loginUserResponse struct {
	User         dbrepo.TbUser `json:"user"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
}

// 公共创建token方法，提供给登录和刷新token
func (s *Server) createToken(uid int64, duration time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "booking sys",
		Subject:   "Booking Apps",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
	}

	// 生成 token
	return s.jwt.GenToken(jwtutil.NewJWTPayload(uid, claims))
}

// loginUser godoc
// @Summary 用户登录
// @Description 手机验证码登录
// @Tags Auth
// @Accept json
// @produce json
// @param json body main.loginUserRequest true "user login param"
// @Router /auth/login [post]
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

	// 验证码登录
	if req.LoginType == LoginTypeCode {
		if s.config.Server.RunMode == "release" {
			// TODO: 真实的短信验证码
			fmt.Println("待实现接入第三方短信接口。")
		} else {
			// 验证短信验证码
			ss, ok := mocksms.GetMockSMS(req.PhoneNumber)
			if !ok || ss.Code() != req.Code {
				app.ToErrorResponse(c, errs.BadRequest.AsMessage("验证码不匹配"))
				return
			}
		}
	} else if req.LoginType == LoginTypePass {
		pass := ""
		if user.Password != nil {
			pass = *user.Password
		}
		if err := pswd.CheckPassword(req.Password, pass); err != nil {
			app.ToErrorResponse(c, errs.BadRequest.AsMessage("密码不匹配"))
			return
		}
	} else {
		app.ToErrorResponse(c, errs.BadRequest.AsMessage("登录方式无效"))
		return
	}

	// 生成 access token
	accessToken, err := s.createToken(user.ID, s.config.JWT.AccessTokenDuration)
	if err != nil {
		app.ToErrorResponse(c, errs.ServerError.AsException(err))
		return
	}

	// 生成 refresh token
	refreshToken, err := s.createToken(user.ID, s.config.JWT.RefreshTokenDuration)
	if err != nil {
		app.ToErrorResponse(c, errs.ServerError.AsException(err))
		return
	}

	// 响应
	response := loginUserResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	app.ToResponse(c, response)
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type refreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (s *Server) refreshToken(c *gin.Context) {
	var req refreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		app.ToErrorResponse(c, errs.BadRequest.AsException(err))
		return
	}

	payload, err := s.jwt.ParseToken(req.RefreshToken)
	if err != nil {
		app.ToErrorResponse(c, errs.Unauthorized.AsException(err, "refresh_token 无效"))
		return
	}

	// 查询用户是否还存在
	_, err = s.store.GetUser(c, payload.UID)
	if err != nil {
		app.ToErrorResponse(c, errs.ServerError.AsException(err, "刷新 token 失败"))
		return
	}

	// 创建 token
	accessToken, err := s.createToken(payload.UID, s.config.JWT.AccessTokenDuration)
	if err != nil {
		app.ToErrorResponse(c, errs.ServerError.AsException(err))
		return
	}

	res := refreshTokenResponse{
		AccessToken: accessToken,
	}

	app.ToResponse(c, res)
}
