package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"toolkit/errs"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/app"
)

func (s *Server) getListRoles(c *gin.Context) {
	roles, err := s.store.GetRoles(context.TODO(), dbrepo.GetRolesParams{Limit: 10, Offset: 0})
	if err != nil {
		// TODO: handler postgreSQL error
		e := errs.InvalidParams.AsException(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, roles)
}

func (s *Server) getRoleById(c *gin.Context) {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err)
		e := errs.InvalidParams.AsException(err, "id invalid")
		app.ToErrorResponse(c, e)
		return
	}
	role, err := s.store.GetRole(context.TODO(), int64(id))
	if err != nil {
		fmt.Println(err)
		e := errs.ServerError.AsException(err)
		app.ToErrorResponse(c, e)
		return
	}
	app.ToResponse(c, role)
}
