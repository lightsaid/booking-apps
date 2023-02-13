package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createUser(c *gin.Context) {
	c.String(http.StatusOK, "Create User.")
}

func (s *Server) updateUser(c *gin.Context) {
	c.String(http.StatusOK, "Update User.")
}

func (s *Server) getListUsers(c *gin.Context) {
	c.String(http.StatusOK, "List User.")
}

func (s *Server) getUserById(c *gin.Context) {
	c.String(http.StatusOK, "Get One User.")
}
