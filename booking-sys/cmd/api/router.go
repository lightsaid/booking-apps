package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	roleRouter := router.Group("/v1/roles")
	{
		roleRouter.GET("", server.getListRoles)
		roleRouter.GET("/:id", server.getRoleById)
		roleRouter.POST("/tx", func(c *gin.Context) {
			id, err := server.store.TestRoleTx(context.TODO())
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"id": id})
		})
	}

	server.router = router
}
