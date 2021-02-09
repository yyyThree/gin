// 404路由
package router

import (
	"fmt"
	"gin/output"
	"github.com/gin-gonic/gin"
)

type noRouter struct {
}

func (noRouter *noRouter) router(router *gin.Engine) {
	fmt.Println("no_router")
	router.NoRoute(func (c *gin.Context) {
		fmt.Println("no_router2")
		output.NotFound(c)
	})
}