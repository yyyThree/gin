// item接口子路由
package group

import (
	"gin/controller"
	"gin/middleware"
	"github.com/gin-gonic/gin"
)

type Item struct {
}

func (item *Item) Router(router *gin.Engine) {
	itemRouter := router.Group("/item")
	itemRouter.Use(new(middleware.Item).Handler())
	itemController := new(controller.Item)
	{
		itemRouter.POST("/add", itemController.Add)
		itemRouter.PUT("/update", itemController.Update)
		itemRouter.DELETE("/delete", itemController.Delete)
		itemRouter.PATCH("/recover", itemController.Recover)
		itemRouter.GET("/get", itemController.Get)
		itemRouter.GET("/search", itemController.Search)
	}
}
