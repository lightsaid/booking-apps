package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"toolkit/errs"
	"toolkit/mocksms"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
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

func (p *pagingRequrest) GetPageNum() int32 {
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	return (p.PageNum - 1) * p.PageSize
}

type idUriRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) pingHandle(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}

// sendSMS godoc
// @Summary 发送短信验证码
// @Description 发送短信验证码，如果是开发模式会直接返回验证码
// @Tags Other
// @Accept json
// @Produce json
// @Param json body main.SMSCode true "手机号"
// @Success 200 {object} any
// @Router /sms [post]
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

func (s *Server) uploadFiles(c *gin.Context) {
	// 控制Body数据大小，包括文件和Form表单其他字段数据，假如想控制文件上传大小不能超过4M,
	// 需要多设置512kb或者1MB给表单其他数据
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, s.config.UploadMaxSize+1024)

	// 上传的文件存储在maxMemory大小的内存里面，如果文件大小超过了maxMemory，
	// 那么剩下的部分将存储在系统的临时文件中。
	err := c.Request.ParseMultipartForm(2 << 20)
	if err != nil {
		app.ToErrorResponse(c, errs.BadRequest.AsException(err, "文件过大"))
		return
	}
	urls := []string{}
	fHeaders := c.Request.MultipartForm.File["files"]
	if len(fHeaders) == 0 {
		log.Println("没有数组")
		app.ToErrorResponse(c, errs.BadRequest.AsMessage("没有文件"))
		return
	}
	for _, file := range fHeaders {
		// 打开文件
		src, err := file.Open()
		if err != nil {
			app.ToErrorResponse(c, errs.BadRequest.AsException(err, "无法打开文件"))
			return
		}
		defer src.Close()

		id, err := uuid.NewV4()
		if err != nil {
			app.ToErrorResponse(c, errs.ServerError.AsException(err))
			return
		}
		ext := filepath.Ext(file.Filename)
		filename := s.config.UploadSavePath + "/" + id.String() + ext
		// 创建文件
		dst, err := os.Create(filename)
		if err != nil {
			app.ToErrorResponse(c, errs.BadRequest.AsException(err, "无法创建文件"))
			return
		}
		defer dst.Close()
		// 复制文件内容
		if _, err = io.Copy(dst, src); err != nil {
			app.ToErrorResponse(c, errs.BadRequest.AsException(err, "无法复制文件内容"))
			return
		}
		urls = append(urls, filename)
	}
	app.ToResponse(c, urls)
}
