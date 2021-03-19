package controller

import (
	"fmt"
	"gin/helper"
	"gin/model/field"
	"gin/model/param"
	"gin/output"
	"gin/output/code"
	"gin/service"
	"github.com/gin-gonic/gin"
)

type Item struct {
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

// 更新商品
func (item *Item) Update(c *gin.Context) {
	params := param.ItemUpdate{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	err := (&service.Item{}).Update(params)
	output.Response(c, nil, err)
	return
}

// 删除商品
func (item *Item) Delete(c *gin.Context) {
	params := param.ItemDelete{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	err := (&service.Item{}).Delete(params)
	output.Response(c, nil, err)
	return
}

// 恢复商品
func (item *Item) Recover(c *gin.Context) {
	params := param.ItemRecover{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	err := (&service.Item{}).Recover(params)
	output.Response(c, nil, err)
	return
}

// 获取商品详情
func (item *Item) Get(c *gin.Context) {
	params := &param.ItemGet{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	data, err := (&service.Item{}).Get(params)
	if err != nil {
		output.Response(c, nil, err)
		return
	}

	fmt.Println("fields", params.Fields, data, field.GetItemFields())
	output.Response(c, &output.SucResponse{
		Data: helper.FilterStructByFields(data, params.Fields, field.GetItemFields()),
	}, nil)
	return
}

// 搜索商品列表 TODO
func (item *Item) Search(c *gin.Context) {
	fmt.Println("ItemDel")
}
