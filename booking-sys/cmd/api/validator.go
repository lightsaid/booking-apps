package main

import (
	"errors"
	"regexp"

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
		// 此时 vphone 是 binding tag 的关键字了
		err = v.RegisterValidation("vphone", vphone)
	} else {
		err = errors.New("validator engine is invalid")
	}

	return err
}
