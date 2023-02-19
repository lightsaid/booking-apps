package app

import (
	"toolkit/errs"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// BindRequest 如果正常绑定返回true；反之处理错误并返回false 并对请求作出响应的错误处理
func BindRequest(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		return bind(c, err)
	}

	return true
}

// BindRequestUri 绑定param参数，如：/api/users/:id 绑定id
func BindRequestUri(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindUri(req); err != nil {
		return bind(c, err)
	}
	return true
}

func bind(c *gin.Context, err error) bool {
	// 获取翻译组件实例
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		// 如果trans经过setTranslations中间件设置后还获取不到，那就是服务内部出问题
		ToErrorResponse(c, errs.ServerError.AsException(err))
		return false
	}
	// 断言错误是否为 validator/v10 的验证错误信息
	verrs, ok := err.(validator.ValidationErrors)
	if !ok { // 其他方面的参数不匹配
		ToErrorResponse(c, errs.InvalidParams.AsException(err))
		return false
	}

	// 对错误信息进行翻译, 得到的是 map[string]string 结构数据
	merrs := verrs.Translate(trans)

	// 拼接错误消息
	var msg string
	var index = 0
	for _, v := range merrs {
		if index > 0 {
			msg += ";"
		}
		msg += v
		index++
	}

	ToErrorResponse(c, errs.InvalidParams.AsException(err, msg))

	return false
}
