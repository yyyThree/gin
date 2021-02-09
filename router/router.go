// 路由中心
package router

import (
	"gin/middleware"
	"github.com/gin-gonic/gin"
)

type routerGroup interface {
	router(*gin.Engine)
}

// 全局中间件 切片
var middlewareHandlers = []func() gin.HandlerFunc{
	new(middleware.Jwt).Handler,
	new(middleware.Panic).Handler,
}

// 子路由群组 切片
var routerGroups = []routerGroup{
	new(noRouter),
	new(itemGroup),
}

// 路由中心入口
func Router() *gin.Engine {

	router := gin.New()

	// 全局中间件注册
	for _, middlewareHandler := range middlewareHandlers {
		router.Use(middlewareHandler())
	}

	// 子路由注册
	for _, routerGroup := range routerGroups {
		register(routerGroup, router)
	}

	return router
}

// 子路由注册
func register(routerGroup routerGroup, router *gin.Engine)  {
	routerGroup.router(router)
}