package controller

import (
	"fmt"
	"gin/helper"
	"gin/model/param"
	"gin/output"
	"gin/output/code"
	"gin/service"
	"github.com/gin-gonic/gin"
)

type Item struct {
}

type ItemAddParams struct {
	Name string `form:"name"`
}

// 添加商品
func (item *Item) Add(c *gin.Context) {
	params := param.ItemAdd{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	data, err := (&service.Item{}).Add(params)
	if err != nil {
		output.Response(c, nil, err)
		return
	}

	output.Response(c, &output.SucResponse{
		Data: data,
	}, nil)
	return
}

func (item *Item) Update(c *gin.Context) {
	fmt.Println("ItemUpdate")
}

func (item *Item) Delete(c *gin.Context) {
	fmt.Println("ItemDel")
}

func (item *Item) Recover(c *gin.Context) {
	fmt.Println("ItemDel")
}

func (item *Item) Get(c *gin.Context) {
	fmt.Println("ItemDel")
}

func (item *Item) Search(c *gin.Context) {
	fmt.Println("ItemDel")
}
