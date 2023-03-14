package errs

import (
	"fmt"
	"net/http"
)

// AppError 一个错误结构体，统一返回错误结构，包含用户错误提示、错误编码和错误信息
type AppError struct {
	message   string
	exception error
	code      uint16
}

// 收集错误状态码和错误信息
var codeMaps = map[uint16]string{}

// 收集 http statusCode
var statusCodeMaps = make(map[uint16]int)

// NewAppError 往 codeMaps 添加错误，如果 code 已存在会触发panic;
// code 是自定义业务编码; msg 是人类可读错误提示; statusCode 是用于 StatusCode 方法返回的http的状态码，有且仅有一个;
// 主要是用于自定义业务码返回 statusCode
func NewAppError(code uint16, msg string, statusCode ...int) *AppError {
	if _, ok := codeMaps[code]; ok {
		panic(fmt.Sprintf("code(%d) already exist, please replace it", code))
	}

	codeMaps[code] = msg

	if len(statusCode) > 0 {
		statusCodeMaps[code] = statusCode[0]
	}

	return &AppError{
		code:    code,
		message: msg,
	}
}

// Error 实现 error 接口
func (a *AppError) Error() string {
	err := fmt.Sprintf("code: %d, message: %s", a.code, a.message)
	if a.exception != nil {
		err = fmt.Sprintf("%s, exception: %v", err, a.exception.Error())
	}
	return err
}

// Unwrap 解开，提供给 errors.Is 和 errors.As 使用
func (a *AppError) Unwrap() error {
	return a.exception
}

// Code 返回错误码
func (a *AppError) Code() uint16 {
	return a.code
}

// Message 返回用户可读错误信息
func (a *AppError) Message() string {
	return a.message
}

// AsMessage 修改消息，例如：错误时入参错误（InvalidParams），可以修改成具体的错误消息（exp:手机号码不正确）
// 返回一个新的 AppError 指针
func (a *AppError) AsMessage(msg string) *AppError {
	return &AppError{
		code:      a.code,
		message:   msg,
		exception: a.exception,
	}
}

// AsException 添加/追加错误, 返回一个新的 AppError 指针
func (a *AppError) AsException(err error, msgs ...string) *AppError {
	var e error
	if a.exception == nil {
		e = fmt.Errorf("%w", err)
	} else {
		e = fmt.Errorf("%v | %w", a.exception, err)
	}
	newErr := &AppError{
		code:      a.code,
		message:   a.message,
		exception: e,
	}
	if len(msgs) > 0 {
		newErr.message = msgs[0]
	}
	return newErr
}

// StatusCode 根据AppError的Code返回 http status code
func (e *AppError) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case Created.Code():
		return http.StatusCreated
	case Accepted.Code():
		return http.StatusAccepted
	case NoContent.Code():
		return http.StatusNoContent
	case BadRequest.Code():
		return http.StatusBadRequest
	case Unauthorized.Code():
		return http.StatusUnauthorized
	case Forbidden.Code():
		return http.StatusForbidden
	case NotFound.Code():
		return http.StatusNotFound
	case MethodNotAllowed.Code():
		return http.StatusMethodNotAllowed
	case NotAcceptable.Code():
		return http.StatusNotAcceptable
	case RequestTimeout.Code():
		return http.StatusRequestTimeout
	case RequestEntityTooLarge.Code():
		return http.StatusRequestEntityTooLarge
	case UnprocessableEntity.Code():
		return http.StatusUnprocessableEntity
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	case ServerError.Code():
		return http.StatusInternalServerError
	case BadGateway.Code():
		return http.StatusBadGateway
	case ServiceUnavailable.Code():
		return http.StatusServiceUnavailable
	case GatewayTimeout.Code():
		return http.StatusGatewayTimeout
	}

	if statusCode, ok := statusCodeMaps[e.code]; ok {
		return statusCode
	}

	return http.StatusInternalServerError
}
