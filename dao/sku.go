package dao

import (
	"gin/constant"
	"gin/model/entity"
	"gorm.io/gorm"
)

type sku struct {
	dao
}

func NewSku(txs ...*gorm.DB) *sku {
	return &sku{dao{
		Tx: GetTx(txs...),
		DbName: constant.DbServiceItems,
	}}
}

func (sku *sku) InsertBatch(insert []*entity.Skus) (data []*entity.Skus, err error) {
	db, err := sku.GetDb()
	if err != nil {
		return
	}

	err = db.Create(&insert).Error
	if err != nil {
		return
	}
	data = insert
	return
}

func (sku *sku) GetList(fields []string, where map[string]interface{}, limit int, offset int) (data []*entity.Skus, err error) {
	db, err := sku.GetDb()
	if err != nil {
		return
	}

	err = db.Select(fields).
		Where(where).
		Limit(limit).
		Offset(offset).
		Find(&data).Error
	return
}

func (sku *sku) Update(update map[string]interface{}, where map[string]interface{}, limit int) (err error) {
	db, err := sku.GetDb()
	if err != nil {
		return
	}
	err = db.Model(&entity.Skus{}).
		Where(where).
		Limit(limit).
		Updates(update).Error
	if err != nil {
		return
	}
	return
}