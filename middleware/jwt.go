package middleware

import (
	"fmt"
	"gin/config"
	"gin/constant"
	"gin/helper"
	tokenLibrary "gin/library/token"
	"gin/output"
	"gin/output/code"
	"github.com/gin-gonic/gin"
	"strings"
)

type Jwt struct {
}

// 请求权限校验
func (jwt *Jwt) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("权限校验中间件开始")

		data, err := checkToken(c)
		if err != nil {
			output.Response(c, nil, err)
			c.Abort()
		}

		c.Set("AppKey", data["appKey"])
		c.Set("Channel", helper.InterfaceToInt(data["channel"]))

		c.Next()

		// after request
		fmt.Println("权限校验中间件结束")
	}
}

// 校验token合法性
func checkToken(c *gin.Context) (data constant.BaseMap, err error) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		err = output.Error(code.NoAuthorization)
		return
	}

	if !strings.HasPrefix(token, "Bearer ") {
		err = output.Error(code.AuthorizationErr)
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)

	tokenStruct := tokenLibrary.New()
	tokenStruct.SetSecret(config.Config.App.TokenSecret)
	tokenStruct.SetToken(token)
	data, errCode := tokenStruct.Decode()

	if errCode > 0 {
		err = output.Error(errCode)
		return
	}

	fmt.Println(data)

	// 校验token必要参数
	if _, ok := data["appKey"]; !ok {
		err = output.Error(code.TokenNotValid)
		return
	}
	if _, ok := data["channel"]; !ok {
		err = output.Error(code.TokenNotValid)
		return
	}

	return
}
