package service

import (
	"gin/constant"
	"gin/entity"
	"gin/helper"
	"gin/model"
	"gin/valid"
)

type Item struct {
	StoreID int
	itemID string
	Name string
	state int
}

func (item *Item) Add() (state int, msgCode string, data constant.BaseMap){
	if err := item.checkAdd(); err != "" {
		state = constant.ParamsValidError
		msgCode = err
		return
	}

	// step1 创建商品ID
	itemId := helper.Uuid()

	// step2 添加商品
	itemModel := &model.Item{
		Item: entity.Item{
			StoreId: 1,
			ItemId: itemId,
			Name: item.Name,
		},
	}
	_, _ = itemModel.Add()

	return
}

// 参数校验
func (item *Item) checkAdd() string {
	valid.New("name", item.Name).IsNotNull()
	if errMsg := valid.Valid(); errMsg != "" {
		return errMsg
	}
	return ""
}