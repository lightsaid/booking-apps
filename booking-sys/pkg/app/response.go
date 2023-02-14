package app

import (
	"log"
	"net/http"
	"toolkit/errs"

	"github.com/gin-gonic/gin"
)

// ToResponse 请求成功响应的处理
func ToResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": errs.Success.Code(),
		"msg":  errs.Success.Message(),
		"data": data,
	})
}

// ToErrorResponse 请求异常的响应处理
func ToErrorResponse(c *gin.Context, err *errs.AppError) {
	// TODO: logger
	log.Printf("metod:%s, url: %s, error: %v", c.Request.Method, c.Request.URL, err)
	if err == nil {
		err = errs.ServerError
	}
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Message(),
	}
	c.JSON(err.StatusCode(), response)
}
