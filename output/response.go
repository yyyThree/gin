package output

import (
	"gin/constant"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type baseRes struct {
	State int    `json:"state"`
	Msg   string `json:"msg"`
}

type sucRes struct {
	baseRes
	Data interface{} `json:"data"`
}

type sucListRes struct {
	sucRes
	Total int `json:"total"`
}

type failRes struct {
	baseRes
}

type errorRes struct {
	baseRes
	Data interface{} `json:"data"`
}

// 通用接口返回
func response(c *gin.Context, httpCode int, res interface{}) {
	c.JSON(httpCode, res)
	return
}

// 成功返回
func Suc(c *gin.Context, data interface{}) {
	res := sucRes{
		baseRes{
			State: constant.ApiSuc,
			Msg:   getMsg("Common.Api." + strconv.Itoa(constant.ApiSuc)),
		},
		data,
	}
	response(c, http.StatusOK, res)
	return
}

// 列表成功返回
func SucList(c *gin.Context, data interface{}, total int) {
	res := sucListRes{
		sucRes{
			baseRes{
				State: constant.ApiSuc,
				Msg:   getMsg("Common.Api." + strconv.Itoa(constant.ApiSuc)),
			},
			data,
		},
		total,
	}
	response(c, http.StatusOK, res)
	return
}

// 失败返回
func Fail(c *gin.Context, state int, msgCode string) {
	switch state {
	case constant.ParamsBindError:
		BindFail(c, msgCode)
	case constant.ParamsValidError:
		ValidFail(c, msgCode)
	default:
		// 获取对应的msg
		msg := getMsg(msgCode)
		res := failRes{
			baseRes{
				State: state,
				Msg:   msg,
			},
		}
		response(c, http.StatusOK, res)
	}
	return
}


// 参数绑定错误
func BindFail(c *gin.Context, msg string) {
	// 获取对应的msg
	res := failRes{
		baseRes{
			State: constant.ParamsBindError,
			Msg:   strings.Join([]string{getMsg("Common.Valid.20001"), msg}, " "),
		},
	}
	response(c, http.StatusBadRequest, res)
	return
}

// 参数校验错误
func ValidFail(c *gin.Context, msg string) {
	// 获取对应的msg
	res := failRes{
		baseRes{
			State: constant.ParamsValidError,
			Msg:   strings.Join([]string{getMsg("Common.Valid.20002"), msg}, " "),
		},
	}
	response(c, http.StatusBadRequest, res)
	return
}

// 请求权限校验失败返回
func AuthFail(c *gin.Context, state int) {
	res := failRes{
		baseRes{
			State: state,
			Msg:   getMsg("Common.Auth." + strconv.Itoa(state)),
		},
	}
	response(c, http.StatusUnauthorized, res)
	return
}

// 请求异常-500
func Error(c *gin.Context, data interface{}) {
	res := errorRes{
		baseRes{
			State: constant.ApiError,
			Msg:   getMsg("Common.Api." + strconv.Itoa(constant.ApiError)),
		},
		data,
	}
	response(c, http.StatusInternalServerError, res)
	return
}

// 404
func NotFound(c *gin.Context) {
	res := baseRes{
		State: constant.ApiNotFound,
		Msg:   getMsg("Common.Api." + strconv.Itoa(constant.ApiNotFound)),
	}
	response(c, http.StatusNotFound, res)
	return
}
