package constant

// 接口标准返回状态码
const (
	ApiSuc      = 1
	ApiNotFound = 404
	ApiError = 500

	TokenValid = 1
	TokenNotFound = 1001
	TokenNotValidPrefix = 1002
	TokenMalformed = 1003
	TokenExpired = 1004
	TokenNotValidYet = 1005
	TokenNotValid = 1006
)
