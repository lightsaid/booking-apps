package main

import (
	"context"
	"net/http"
	"toolkit/dberr"

	"github.com/gin-gonic/gin"
	"github.com/lightsaid/booking-sys/pkg/app"
)

func (s *Server) initRouter() {
	router := gin.Default()

	router.Use(s.setTranslations())

	// 服务检查
	router.GET("/v1/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})

	if s.config.Server.RunMode == "release" {
		// TODO: 真实发送短信验证码
		router.POST("/v1/sms")
	} else {
		// 模拟发送短信验证码
		router.POST("/v1/sms", s.mockSendSMS)
	}

	authRouter := router.Group("/v1/auth")
	{
		authRouter.POST("/login", s.loginUser)
		authRouter.POST("/refresh", s.refreshToken)
	}

	userRouter := router.Group("/v1/users").Use(s.authentication())
	{
		userRouter.POST("", s.createUser)
		userRouter.PUT("/:id", s.updateUser)
		userRouter.GET("", s.getListUsers) // /v1/users?page_num=1&page_size=10
		userRouter.GET("/:id", s.getUserById)
	}

	roleRouter := router.Group("/v1/roles").Use(s.authentication())
	{
		roleRouter.GET("", s.getListRoles)
		roleRouter.GET("/:id", s.getRoleById)
		roleRouter.POST("/tx", func(c *gin.Context) {
			id, err := s.store.TestRoleTx(context.TODO())
			if err != nil {
				e, _ := dberr.HandleDBError(err)
				app.ToErrorResponse(c, e)
				return
			}
			c.JSON(http.StatusOK, gin.H{"id": id})
		})
	}

	s.router = router
}
