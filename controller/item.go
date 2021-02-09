package controller

import (
	"fmt"
	"gin/output"
	"github.com/gin-gonic/gin"
)

type Item struct {
}

func (item *Item) Get(c *gin.Context) {
	itemId := c.Param("itemId")
	fields := c.DefaultQuery("fields", "")
	fmt.Println("ItemGet", itemId, fields)
	output.Suc(c, map[int]string{0: "test"})
	output.SucList(c, map[int]string{0: "test"}, 2)
	output.Fail(c, 2001, "Item.Get.2001")
	return
}

func (item *Item) Add(c *gin.Context) {
	storeId := c.PostForm("storeId")
	itemId := c.PostForm("itemId")
	name := c.PostForm("name")
	test := c.PostFormMap("test")
	fmt.Println("ItemAdd", storeId, itemId, name, test)
}

func (item *Item) Update(c *gin.Context) {
	fmt.Println("ItemUpdate")
}

func (item *Item) Del(c *gin.Context) {
	fmt.Println("ItemDel")
}
