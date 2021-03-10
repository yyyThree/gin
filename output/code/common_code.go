package code

//go:generate stringer -type Code -linecomment -output common_string.go

const (
	OK          Code = 1   // 成功
	ApiNotFound Code = 404 // 接口不存在

	NoAuthorization  Code = 1001 // 未获取到Authorization
	AuthorizationErr Code = 1002 // Authorization非法
	TokenNotFound    Code = 1003 // token不存在
	TokenMalformed   Code = 1004 // token解析失败
	TokenExpired     Code = 1005 // token已失效
	TokenNotValidYet Code = 1006 // token未生效
	TokenNotValid    Code = 1007 // token无效

	ParamBindErr  Code = 2001 // 参数绑定失败
	IllegalParams Code = 2002 // 参数非法

	RecordNotFound Code = 3001 // 未查询到记录

	MySqlErr Code = 4001 // mysql执行错误

	ServerErr Code = 5001 // 服务器错误
)
