package main

import (
	"fmt"
	"time"
	"toolkit/mocksms"

	"github.com/gin-gonic/gin"
	"github.com/lightsaid/booking-sys/pkg/app"
)

type SMSCode struct {
	PhoneNumber string `json:"phone_number" zh:"手机号码" binding:"required,len=11"`
}

func (s *Server) mockSendSMS(c *gin.Context) {
	var req SMSCode
	if ok := app.BindRequest(c, &req); !ok {
		return
	}

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

type PagingRequrest struct {
	// 每页多少条数据
	PageSize int32 `form:"page_size" binding:"required,min=10,max=100"`
	// 第几页
	PageNum int32 `form:"page_num" binding:"required,min=1"`
}
