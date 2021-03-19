// 路由中心
package router

import (
	"gin/middleware"
	"gin/router/group"
	"github.com/gin-gonic/gin"
)

type Group interface {
	Router(*gin.Engine)
}

type MiddleWare interface {
	Handler() gin.HandlerFunc
}

// 全局中间件 切片
var globalMiddlewares = []MiddleWare{
	new(middleware.Panic),
	new(middleware.Jwt),
}

// 子路由群组 切片
var routerGroups = []Group{
	new(group.NoRouter),
	new(group.Item),
}

// 路由中心入口
func Router() *gin.Engine {

	router := gin.New()

	// 全局中间件注册
	for _, globalMiddleware := range globalMiddlewares {
		router.Use(globalMiddleware.Handler())
	}

	// 子路由注册
	for _, routerGroup := range routerGroups {
		routerGroup.Router(router)
	}

	return router
}
