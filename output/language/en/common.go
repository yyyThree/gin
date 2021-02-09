package en

var Common = map[string]map[int]string{
	"Api": {
		1:   "成功",
		404: "接口不存在",
	},
	"Auth": {
		1001: "token不存在",
		1002: "token非法",
		1003: "token校验失败",
	},
}
