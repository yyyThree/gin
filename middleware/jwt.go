package middleware

import (
	"fmt"
	"gin/config"
	"gin/constant"
	tokenLibrary "gin/library/token"
	"gin/output"
	"github.com/gin-gonic/gin"
	"strings"
)

type TokenData struct {
	storeId int
}

type Jwt struct {
}

// 请求权限校验
func (jwt *Jwt) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("权限校验中间件开始")

		data, state := checkToken(c)
		if state != 1 {
			output.AuthFail(c, state)
			c.Abort()
		}

		c.Set("StoreId", data["store_id"])

		// before request

		c.Next()

		// after request
		fmt.Println("权限校验中间件结束")
	}
}

// 校验token合法性
func checkToken(c *gin.Context) (data constant.BaseMap, state int) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		state = constant.TokenNotFound
		return
	}

	if !strings.HasPrefix(token, "Bearer ") {
		state = constant.TokenNotValidPrefix
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)

	tokenStruct := tokenLibrary.New()
	tokenStruct.SetSecret(config.Config.App.TokenSecret)
	tokenStruct.SetToken(token)
	data, state = tokenStruct.Decode()

	fmt.Println(data)
	// 校验token必要参数
	if _, ok := data["store_id"]; !ok {
		state = constant.TokenParamsNotValid
		return
	}

	return
}
