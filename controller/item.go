package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yyyThree/gin/helper"
	"github.com/yyyThree/gin/model/field"
	"github.com/yyyThree/gin/model/param"
	"github.com/yyyThree/gin/output"
	"github.com/yyyThree/gin/output/code"
	"github.com/yyyThree/gin/service"
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

	output.Response(c, &output.SucResponse{
		Data: helper.FilterStructByFields(data, params.Fields, field.GetItemFields()),
	}, nil)
	return
}

// 搜索商品列表
func (item *Item) Search(c *gin.Context) {
	params := &param.ItemSearch{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	data, total, err := (&service.ItemSearch{}).Search(params)
	if err != nil {
		output.Response(c, nil, err)
		return
	}

	output.Response(c, &output.ListResponse{
		SucResponse: &output.SucResponse{
			Data: helper.FilterStructsByFields(data, params.Fields, field.GetItemFields()),
		},
		Total: total,
	}, nil)
	return
}
