package zh

import "gin/constant"

var Common = constant.MsgMap{
	"Api": {
		1:   "成功",
		404: "接口不存在",
		500: "请求异常",
	},
	"Auth": {
		1001: "token不存在",
		1002: "token非法",
		1003: "token格式不正确",
		1004: "token已过期",
		1005: "token尚未生效",
		1006: "token校验失败",
		1007: "token参数缺失",
	},
	"Valid": {
		20001: "参数格式错误",
		20002: "参数校验错误",
	},
}
