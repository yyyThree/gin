// 404路由
package router

import (
	"gin/output"
	"gin/output/code"
	"github.com/gin-gonic/gin"
)

type noRouter struct {
}

func (noRouter *noRouter) router(router *gin.Engine) {
	router.NoRoute(func (c *gin.Context) {
		output.Response(c, nil, output.Error(code.ApiNotFound))
	})
}