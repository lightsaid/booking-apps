package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
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

	userRouter := router.Group("/v1/users")
	{
		userRouter.POST("", s.createUser)
		userRouter.PUT("/:id", s.updateUser)
		userRouter.GET("", s.getListUsers)
		userRouter.GET("/:id", s.getUserById)
	}

	roleRouter := router.Group("/v1/roles")
	{
		roleRouter.GET("", s.getListRoles)
		roleRouter.GET("/:id", s.getRoleById)
		roleRouter.POST("/tx", func(c *gin.Context) {
			id, err := s.store.TestRoleTx(context.TODO())
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"id": id})
		})
	}

	s.router = router
}
