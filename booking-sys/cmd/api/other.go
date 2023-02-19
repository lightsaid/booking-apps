package main

import (
	"fmt"
	"net/http"
	"time"
	"toolkit/mocksms"

	"github.com/gin-gonic/gin"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type SMSCode struct {
	PhoneNumber string `json:"phone_number" zh:"手机号码" binding:"required,len=11"`
}

type pagingRequrest struct {
	// 每页多少条数据
	PageSize int32 `form:"page_size" binding:"required,min=10,max=100"`
	// 第几页
	PageNum int32 `form:"page_num" binding:"required,min=1"`
}

type idUriRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) pingHandle(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}

func (s *Server) sendSMS(c *gin.Context) {
	var req SMSCode
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

	if s.config.Server.RunMode == "release" {
		// TODO: 真实发送短信验证码

	} else {
		// 模拟发送短信验证码
		sms := mocksms.NewMockSMS(req.PhoneNumber)
		go func() {
			time.Sleep(3 * time.Second)
			sms.SetStatus(req.PhoneNumber, mocksms.StatusOpts.Expired)
			fmt.Println(sms.Code(), sms.Status())

			v, _ := mocksms.GetMockSMS(req.PhoneNumber)
			fmt.Println("get: ", v)
		}()
		app.ToResponse(c, sms.Code())
	}
}
