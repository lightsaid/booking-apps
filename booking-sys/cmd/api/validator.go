package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	PhoneNumberRx = regexp.MustCompile(`^1[3-9]{1}\d{9}$`)
)

// vphone 自定义验证11位手机号码函数
var vphone validator.Func = func(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return PhoneNumberRx.MatchString(value)
	}
	return false
}

func setupValidatorEngine() error {
	var err error
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取字段tag的自定义方法, 分别获取 json tag 和 zh tag, 提供给错误消息使用
		v.RegisterTagNameFunc(func(fl reflect.StructField) string {
			fmt.Println("fl json : ", fl.Tag.Get("json"))
			fmt.Println("fl zh: ", fl.Tag.Get("zh"))

			name := fl.Tag.Get("zh")
			if len(strings.TrimSpace(name)) == 0 {
				name = fl.Tag.Get("json")
			}
			name += " "
			return name
		})

		// 将 vphone 注册成 binding tag 的关键字
		err = v.RegisterValidation("vphone", vphone)
	} else {
		err = errors.New("validator engine is invalid")
	}

	return err
}
