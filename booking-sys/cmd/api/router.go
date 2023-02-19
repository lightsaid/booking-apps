package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) initRouter() {
	router := gin.Default()
	router.Use(s.setTranslations())

	// 服务检查
	router.GET("/v1/api/ping", s.pingHandle)

	// 发送短信
	router.POST("/v1/api/sms", s.sendSMS)

	// auth 登录/认证模块
	authRouter := router.Group("/v1/api/auth")
	{
		authRouter.POST("/login", s.loginUser)
		authRouter.POST("/refresh", s.refreshToken)
	}

	// admin 管理员模块
	adminRouter := router.Group("/v1/api/admin").Use(s.authentication())
	{
		adminRouter.POST("/users", s.createUser)
		adminRouter.GET("/users", s.getListUsers) // /v1/users?page_num=1&page_size=10
		adminRouter.PUT("/users/:id", s.updateUser)
		adminRouter.GET("/users/:id", s.getUserById)

		adminRouter.GET("/roles", s.getListRoles)
		adminRouter.GET("/roles/:id", s.getRoleById)

		adminRouter.POST("/theaters", s.createTheater)
		adminRouter.GET("/theaters", s.listTheaters)
		adminRouter.GET("/theaters/:id", s.getTheater)
		adminRouter.PUT("/theaters/:id", s.updateTheater)
	}

	s.router = router
}
