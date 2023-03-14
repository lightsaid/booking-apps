package errs

import "net/http"

/**
 * 此文件定义系统的所有错误码, 包括公共错误码和具体业务错误码
 */

var (
	// 定义公共错误码
	Success               = NewAppError(200, "请求成功")    // 表示请求成功，并返回所请求的数据。
	Created               = NewAppError(201, "请求成功")    // 表示请求成功，并创建了新的资源。通常用于 POST 请求。
	Accepted              = NewAppError(202, "请求已接受")   // 示请求已接受，但尚未处理完成。通常用于异步处理的场景，例如批量操作、数据导入等。
	NoContent             = NewAppError(204, "请求成功")    // 表示请求成功，但没有返回任何数据。通常用于 DELETE 请求。
	BadRequest            = NewAppError(400, "入参错误")    //  表示客户端发送的请求有误，服务器无法理解。
	Unauthorized          = NewAppError(401, "验证失败")    // 表示客户端未经身份验证或身份验证失败。
	Forbidden             = NewAppError(403, "未经授权")    // 表示客户端未经授权访问资源。
	NotFound              = NewAppError(404, "未找到")     // 表示请求的资源不存在
	MethodNotAllowed      = NewAppError(405, "请求方法不支持") // 表示请求方法不支持。
	NotAcceptable         = NewAppError(406, "请求头无效")   // 表示请求的 Accept 头部中指定的格式无法被服务器接受。
	RequestTimeout        = NewAppError(408, "请求超时")    // 表示客户端请求超时。
	RequestEntityTooLarge = NewAppError(413, "请求体过大")   // 示客户端请求体太大，服务器无法处理。
	UnprocessableEntity   = NewAppError(422, "入参有误")    // 表示客户端发送的请求格式正确，但服务器无法处理，通常是由于请求体中缺少必要的字段或参数错误。
	TooManyRequests       = NewAppError(429, "请求繁忙")    // 表示客户端发送的请求过多，超出了限制。
	ServerError           = NewAppError(500, "务器内部错误")  // 表示服务器内部错误。
	BadGateway            = NewAppError(502, "网关错误")    // 表示服务器作为网关或代理时，无法访问上游服务器。
	ServiceUnavailable    = NewAppError(503, "无法处理请求")  // 表示服务器暂时无法处理请求，通常是由于服务器过载或正在进行维护。
	GatewayTimeout        = NewAppError(504, "服务器未响应")  // 表示服务器作为网关或代理时，上游服务器超时未响应。

	// 具体业务错误码，以Err开头，必须指定 HTTP StatusCode, 否则按http.StatusInternalServerError处理
	ErrAlreadyExists = NewAppError(10000, "记录已存在", http.StatusBadRequest)
	ErrUserExist     = NewAppError(10001, "用户已存在", http.StatusBadRequest)
)
