// 404路由
package group

import (
	"github.com/gin-gonic/gin"
	"github.com/yyyThree/gin/output"
	"github.com/yyyThree/gin/output/code"
)

type NoRouter struct {
}

func (noRouter *NoRouter) Router(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		output.Response(c, nil, output.Error(code.ApiNotFound))
	})
}
