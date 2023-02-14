package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) loginUser(c *gin.Context) {
	c.String(http.StatusOK, "Login User.")
}
func (s *Server) refreshToken(c *gin.Context) {
	c.String(http.StatusOK, "Refresh Token.")
}
