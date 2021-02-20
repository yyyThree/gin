// 404路由
package router

import (
	"gin/output"
	"github.com/gin-gonic/gin"
)

type noRouter struct {
}

func (noRouter *noRouter) router(router *gin.Engine) {
	router.NoRoute(func (c *gin.Context) {
		output.NotFound(c)
	})
}