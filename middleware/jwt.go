package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yyyThree/gin/config"
	"github.com/yyyThree/gin/constant"
	"github.com/yyyThree/gin/helper"
	tokenLibrary "github.com/yyyThree/gin/library/token"
	"github.com/yyyThree/gin/output"
	"github.com/yyyThree/gin/output/code"
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
	data, err = tokenStruct.Decode()

	if err != nil {
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
