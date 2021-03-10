// item接口子路由
package router

import (
	"gin/controller"
	"gin/middleware"
	"github.com/gin-gonic/gin"
)

type itemGroup struct {
}

func (itemGroup *itemGroup) router(router *gin.Engine) {
	itemRouter := router.Group("/item")
	itemRouter.Use(new(middleware.Item).Handler())
	itemController := new(controller.Item)
	{
		itemRouter.POST("/add", itemController.Add)
		itemRouter.PUT("/update", itemController.Update)
		itemRouter.DELETE("/delete", itemController.Delete)
		itemRouter.PUT("/recover", itemController.Recover)
		itemRouter.GET("/get", itemController.Get)
		itemRouter.GET("/search", itemController.Search)
	}
}