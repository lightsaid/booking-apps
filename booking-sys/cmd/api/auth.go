package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) loginUser(c *gin.Context) {
	c.String(http.StatusOK, "Login User.")
}
func (server *Server) refreshToken(c *gin.Context) {
	c.String(http.StatusOK, "Refresh Token.")
}
