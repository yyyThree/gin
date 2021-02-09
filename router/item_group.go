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
		itemRouter.GET("/get/:itemId", itemController.Get)
		itemRouter.POST("/add", itemController.Add)
		itemRouter.PUT("/update/:itemId", itemController.Update)
		itemRouter.DELETE("/del/:itemId", itemController.Del)
	}
}