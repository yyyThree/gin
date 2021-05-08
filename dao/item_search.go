package dao

import (
	"github.com/yyyThree/gin/constant"
	"github.com/yyyThree/gin/helper"
	"github.com/yyyThree/gin/model/entity"
	"gorm.io/gorm"
)

type itemSearch struct {
	dao
}

func NewItemSearch(txs ...*gorm.DB) *itemSearch {
	return &itemSearch{dao{
		Tx:     GetTx(txs...),
		DbName: constant.DbServiceItems,
	}}
}

func (itemSearch *itemSearch) Insert(insert *entity.ItemSearches) (data entity.ItemSearches, err error) {
	db, err := itemSearch.GetDb()
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

func (itemSearch *itemSearch) GetOne(fields []string, where map[string]interface{}) (data entity.ItemSearches, err error) {
	db, err := itemSearch.GetDb()
	if err != nil {
		return
	}
	err = db.Select(fields).
		Where(where).
		Limit(1).
		Find(&data).Error
	return
}

func (itemSearch *itemSearch) Update(update map[string]interface{}, where map[string]interface{}, limit int) (err error) {
	db, err := itemSearch.GetDb()
	if err != nil {
		return
	}
	err = db.Model(&entity.ItemSearches{}).
		Where(where).
		Limit(limit).
		Updates(update).Error
	if err != nil {
		return
	}
	return
}

func (itemSearch *itemSearch) GetList(fields []string, where map[string]interface{}, limit int, offset int) (data []*entity.ItemSearches, err error) {
	db, err := itemSearch.GetDb()
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

func (itemSearch *itemSearch) InsertBatch(insert []*entity.ItemSearches) (data []*entity.ItemSearches, err error) {
	db, err := itemSearch.GetDb()
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

func (itemSearch *itemSearch) UpdateBatch(update []*entity.ItemSearches) (err error) {
	db, err := itemSearch.GetDb()
	if err != nil {
		return
	}
	err = db.Save(update).Error
	if err != nil {
		return
	}
	return
}

func (itemSearch *itemSearch) SearchItem(sqlBuild constant.SqlBuild) (data []string, total int, err error) {
	db, err := itemSearch.GetDb()
	if err != nil {
		return
	}

	var queryData []entity.ItemSearches
	var count int64
	sql := db.Model(&entity.ItemSearches{}).
		Distinct("item_id").
		Where(sqlBuild.Where)
	itemSearch.Like(sql, sqlBuild.Like)

	sql.Count(&count)
	total = int(count)

	err = sql.
		Limit(sqlBuild.Limit).
		Offset(sqlBuild.Offset).
		Find(&queryData).Error

	if !helper.IsEmpty(queryData) {
		for _, v := range queryData {
			data = append(data, v.ItemID)
		}
	}

	return
}
