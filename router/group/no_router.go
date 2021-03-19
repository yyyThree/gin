// 404路由
package group

import (
	"gin/output"
	"gin/output/code"
	"github.com/gin-gonic/gin"
)

type NoRouter struct {
}

func (noRouter *NoRouter) Router(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		output.Response(c, nil, output.Error(code.ApiNotFound))
	})
}
