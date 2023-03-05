package main

import (
	"net/http"
	"strings"
	"toolkit/errs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/validator/v10"
	"github.com/lightsaid/booking-sys/pkg/app"

	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义常量
const (
	AuthorizationKey        = "Authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

// setTranslations 设置 validator/v10 错误消息的翻译
func (s *Server) setTranslations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}

// authentication 认证用户是否登录
func (s *Server) authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(AuthorizationKey)
		if len(authorizationHeader) == 0 {
			app.ToErrorResponse(c, errs.NeedToLogin)
			c.Abort()
			return
		}

		// 提取 accessToken "Bearer eyJhbGciOiJIUzI...."

		// 以空格分割为两部分
		parts := strings.Fields(authorizationHeader)
		if len(parts) < 2 {
			app.ToErrorResponse(c, errs.UnauthorizedTokenError.AsMessage("Token 格式不匹配"))
			c.Abort()
			return
		}

		// 验证accessToken 头
		authorizationType := strings.ToLower(parts[0])
		if authorizationType != AuthorizationTypeBearer {
			app.ToErrorResponse(c, errs.UnauthorizedTokenError.AsMessage("Token 类型不匹配"))
			c.Abort()
			return
		}

		// 验证accessToken
		accessToken := parts[1]
		payload, err := s.jwt.ParseToken(accessToken)
		if err != nil {
			app.ToErrorResponse(c, errs.UnauthorizedTokenError.AsException(err, err.Error()))
			c.Abort()
			return
		}

		// 设置上下文
		c.Set(AuthorizationPayloadKey, payload)

		c.Next()
	}
}

func (s *Server) setCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
