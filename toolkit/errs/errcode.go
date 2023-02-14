package errs

// 定义公共错误码
var (
	Success                  = NewAppError(1000, "成功")
	InvalidParams            = NewAppError(1001, "参数错误")
	NotFound                 = NewAppError(1002, "未找到")
	NeedToLogin              = NewAppError(1003, "请先登录")
	ServerError              = NewAppError(1004, "服务内部错误")
	NotForbidden             = NewAppError(1005, "权限不足")
	RequestTimeout           = NewAppError(1006, "请求超时")
	UnauthorizedTokenError   = NewAppError(1007, "Token错误")
	UnauthorizedTokenTimeout = NewAppError(1008, "Token超时")
	TooManyRequests          = NewAppError(1009, "请求频繁")
	UnprocessableEntity      = NewAppError(1010, "请求无法处理")
)
