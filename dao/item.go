package dao

import (
	"github.com/yyyThree/gin/constant"
	"github.com/yyyThree/gin/model/entity"
	"gorm.io/gorm"
)

type item struct {
	dao
}

func NewItem(txs ...*gorm.DB) *item {
	return &item{dao{
		Tx:     GetTx(txs...),
		DbName: constant.DbServiceItems,
	}}
}

func (item *item) Insert(insert *entity.Items) (data entity.Items, err error) {
	db, err := item.GetDb()
	if err != nil {
		return
	}

	err = db.Create(&insert).Error
	if err != nil {
		return
	}
	data = *insert
	return
}

func (item *item) GetOne(fields []string, where map[string]interface{}) (data entity.Items, err error) {
	db, err := item.GetDb()
	if err != nil {
		return
	}
	err = db.Select(fields).
		Where(where).
		Limit(1).
		Find(&data).Error
	return
}

func (item *item) Update(update map[string]interface{}, where map[string]interface{}, limit int) (err error) {
	db, err := item.GetDb()
	if err != nil {
		return
	}
	err = db.Model(&entity.Items{}).
		Where(where).
		Limit(limit).
		Updates(update).Error
	if err != nil {
		return
	}
	return
}
