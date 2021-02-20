package model

import (
	"fmt"
	"gin/constant"
	"gin/entity"
	"gin/helper"
)

type Item struct {
	entity.Item
}

func (item *Item) Add() (int, error) {
	if helper.HasAnyEmpty(item.StoreId, item.ItemId, item.Name) {
		return constant.ModelParamsError, nil
	}
	res := GetMasterDB(constant.DbServiceItems).Create(item)

	if res.RowsAffected == 0 {
		return constant.ModelSqlExecError, nil
	}
	fmt.Println(123, res.Error, res.Statement.SQL)

	return constant.ModelSuc, nil
}
