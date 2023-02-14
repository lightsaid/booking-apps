package app

import (
	"github.com/gin-gonic/gin"
)

// BindRequest 如果正常绑定返回true，反之处理错误并返回false
func BindRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		// TODO:
		return false
	}

	return true
}
