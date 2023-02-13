package main

import "github.com/gin-gonic/gin"

func (server *Server) initRouter() {
	router := gin.Default()

	authRouter := router.Group("/v1/auth")
	{
		authRouter.POST("/login", server.loginUser)
		authRouter.POST("/refresh", server.refreshToken)
	}

	userRouter := router.Group("/v1/users")
	{
		userRouter.POST("/", server.createUser)
		userRouter.PUT("/:id", server.updateUser)
		userRouter.GET("/", server.getListUsers)
		userRouter.GET("/:id", server.getUserById)
	}

	server.router = router
}
