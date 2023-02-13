package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createUser(c *gin.Context) {
	c.String(http.StatusOK, "Create User.")
}

func (server *Server) updateUser(c *gin.Context) {
	c.String(http.StatusOK, "Update User.")
}

func (server *Server) getListUsers(c *gin.Context) {
	c.String(http.StatusOK, "List User.")
}

func (server *Server) getUserById(c *gin.Context) {
	c.String(http.StatusOK, "Get One User.")
}
