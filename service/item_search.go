package service

import (
	"github.com/yyyThree/gin/constant"
	"github.com/yyyThree/gin/dao"
	"github.com/yyyThree/gin/helper"
	"github.com/yyyThree/gin/library/valid"
	"github.com/yyyThree/gin/model/entity"
	"github.com/yyyThree/gin/model/param"
	"github.com/yyyThree/gin/output"
	"github.com/yyyThree/gin/output/code"
)

type ItemSearch struct {
}

// 搜索商品列表
func (itemSearch *ItemSearch) Search(params *param.ItemSearch) (data []ItemDetail, total int, err error) {
	if err = itemSearch.checkSearch(params); err != nil {
		return
	}

	// step1 初始化默认值
	params.Fields = helper.GetString(params.Fields, "*")

	// step2 搜索item_id
	// 构建搜索条件
	sqlBuild := itemSearch.buildSearch(params)
	itemIds, total, _ := dao.NewItemSearch().SearchItem(sqlBuild)

	if helper.IsEmpty(itemIds) {
		return
	}

	// step3 根据itemIds获取商品详情
	for _, itemId := range itemIds {
		item, _ := (&Item{}).Get(&param.ItemGet{
			ItemId: itemId,
			Fields: params.Fields,
			Common: params.Common,
		})
		if !helper.IsEmpty(item) {
			data = append(data, item)
		}
	}

	return
}

func (itemSearch *ItemSearch) buildSearch(params *param.ItemSearch) (sqlBuild constant.SqlBuild) {
	// 初始化默认值
	params.Page = helper.GetInt(params.Page, constant.Page)
	params.Limit = helper.GetInt(params.Limit, constant.Limit)

	where := constant.BaseMap{
		"appkey":  params.AppKey,
		"channel": params.Channel,
	}
	like := constant.BaseMap{}
	if !helper.IsEmpty(params.ItemId) {
		where["item_id"] = params.ItemId
	}
	if !helper.IsEmpty(params.ItemName) {
		like["item_name"] = params.ItemName
	}
	if !helper.IsEmpty(params.SkuName) {
		like["sku_name"] = params.SkuName
	}
	if !helper.IsEmpty(params.Barcode) {
		like["barcode"] = params.Barcode
	}
	if !helper.IsEmpty(params.ItemState) {
		where["item_state"] = *params.ItemState
		where["sku_state"] = *params.ItemState
	}

	sqlBuild.Where = where
	sqlBuild.Like = like
	sqlBuild.Limit = params.Limit
	sqlBuild.Offset = (params.Page - 1) * params.Limit
	return
}

func (itemSearch *ItemSearch) Sync(params param.ItemSync) (err error) {
	if err = itemSearch.checkSync(params); err != nil {
		return
	}

	// step1 获取商品详情
	item, err := (&Item{}).Get(&param.ItemGet{
		ItemId: params.ItemId,
		Common: params.Common,
	})
	if err != nil {
		err = output.Error(code.ItemNoFound)
		return
	}

	// step2 根据不同状态进行同步
	switch params.SyncType {
	case constant.ItemSyncTypeAdd: // 添加商品
		err = itemSearch.SyncAdd(item)
	case constant.ItemSyncTypeUpdate, constant.ItemSyncTypeDelete, constant.ItemSyncTypeRecover: // 更新/删除/恢复商品
		err = itemSearch.SyncUpdate(item)
	}
	return
}

// 同步添加商品搜索数据
func (itemSearch *ItemSearch) SyncAdd(item ItemDetail) (err error) {
	// step1 获取商品搜索数据
	itemSearches, _ := dao.NewItemSearch().GetList([]string{}, constant.BaseMap{
		"appkey":  item.Appkey,
		"channel": item.Channel,
		"item_id": item.ItemID,
	}, constant.CommonLimit, 0)

	if !helper.IsEmpty(itemSearches) {
		return itemSearch.SyncUpdate(item)
	}

	// step2 添加商品搜索数据
	var insertBatch []*entity.ItemSearches
	for _, sku := range item.Skus {
		insertBatch = append(insertBatch, &entity.ItemSearches{
			Appkey:    item.Appkey,
			Channel:   item.Channel,
			ItemID:    item.ItemID,
			SkuID:     sku.SkuID,
			ItemName:  item.Name,
			SkuName:   sku.Name,
			Barcode:   sku.Barcode,
			ItemState: item.State,
			SkuState:  sku.State,
		})
	}
	_, err = dao.NewItemSearch().InsertBatch(insertBatch)
	if err != nil {
		err = output.Error(code.ItemSearchInsertFail)
		return
	}

	return
}

// 同步更新商品搜索数据
func (itemSearch *ItemSearch) SyncUpdate(item ItemDetail) (err error) {
	// step1 获取商品搜索数据
	itemSearches, _ := dao.NewItemSearch().GetList([]string{}, constant.BaseMap{
		"appkey":  item.Appkey,
		"channel": item.Channel,
		"item_id": item.ItemID,
	}, constant.CommonLimit, 0)

	if helper.IsEmpty(itemSearches) {
		return itemSearch.SyncAdd(item)
	}

	itemSearchesMap := make(map[string]*entity.ItemSearches)
	for _, itemSearch := range itemSearches {
		itemSearchesMap[itemSearch.SkuID] = itemSearch
	}

	// step2 更新
	var insertBatch []*entity.ItemSearches
	var updateBatch []*entity.ItemSearches
	for _, sku := range item.Skus {
		if itemSearch, ok := itemSearchesMap[sku.SkuID]; ok {
			updateBatch = append(updateBatch, &entity.ItemSearches{
				ID:        itemSearch.ID,
				Appkey:    itemSearch.Appkey,
				Channel:   itemSearch.Channel,
				ItemID:    itemSearch.ItemID,
				SkuID:     sku.SkuID,
				ItemName:  item.Name,
				SkuName:   sku.Name,
				Barcode:   sku.Barcode,
				ItemState: item.State,
				SkuState:  sku.State,
			})
		} else {
			insertBatch = append(insertBatch, &entity.ItemSearches{
				Appkey:    item.Appkey,
				Channel:   item.Channel,
				ItemID:    item.ItemID,
				SkuID:     sku.SkuID,
				ItemName:  item.Name,
				SkuName:   sku.Name,
				Barcode:   sku.Barcode,
				ItemState: item.State,
				SkuState:  sku.State,
			})
		}
	}

	if !helper.IsEmpty(insertBatch) {
		_, err = dao.NewItemSearch().InsertBatch(insertBatch)
		if err != nil {
			err = output.Error(code.ItemSearchInsertFail)
			return
		}
	}

	if !helper.IsEmpty(updateBatch) {
		err = dao.NewItemSearch().UpdateBatch(updateBatch)
		if err != nil {
			err = output.Error(code.ItemSearchUpdateFail)
			return
		}
	}

	return
}

func (itemSearch *ItemSearch) checkSearch(params *param.ItemSearch) error {
	myValid := valid.New()
	if !helper.IsEmpty(params.ItemState) {
		myValid.Append(valid.One("item_state", *params.ItemState).IsIn([]int{constant.ItemStateNormal, constant.ItemStateDeleted, constant.ItemStateFinalDeleted}))
	}

	if err := myValid.Valid(); err != "" {
		return output.Error(code.IllegalParams).WithDetails(err)
	}
	return nil
}

func (itemSearch *ItemSearch) checkSync(params param.ItemSync) error {
	myValid := valid.New()
	myValid.Append(valid.One("item_id", params.ItemId).NotEmpty().MaxStringLength(64))
	myValid.Append(valid.One("sync_type", params.SyncType).NotEmpty().IsIn(constant.ItemSyncTypes))

	if err := myValid.Valid(); err != "" {
		return output.Error(code.IllegalParams).WithDetails(err)
	}
	return nil
}
