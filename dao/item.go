package dao

import (
	"gin/constant"
	"gin/model/db"
	"gin/model/entity"
)

type item struct {
	dao
}

func NewItem() *item{
	return &item{dao{
		db.GetMasterDB(constant.DbServiceItems),
	}}
}

func (item *item) Insert(insert *entity.Items) (data *entity.Items, err error){
	if item.DB == nil {
		return
	}

	err = item.DB.Create(insert).Error
	if err != nil {
		return
	}

	data = insert
	return
}