package controller

import (
	"fmt"
	"gin/constant"
	"gin/output"
	"gin/service"
	"github.com/gin-gonic/gin"
)

type Item struct {
}

type ItemAddParams struct {
	Name string `form:"name"`
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
	params := ItemAddParams{}
	if err := c.ShouldBind(&params); err != nil {
		output.BindFail(c, err.Error())
		return
	}

	// storeId, _ := c.Get("storeId")
	itemService := &service.Item{
		// StoreID: strconv.Atoi(storeId),
		Name: params.Name,
	}
	state, msgCode, data := itemService.Add()
	if state != constant.ApiSuc {
		output.Fail(c, state, msgCode)
		return
	}
	output.Suc(c, data)
	return
}

func (item *Item) Update(c *gin.Context) {
	fmt.Println("ItemUpdate")
}

func (item *Item) Del(c *gin.Context) {
	fmt.Println("ItemDel")
}
