package param

import "github.com/yyyThree/gin/model/entity"

type ItemAdd struct {
	Name   string         `form:"name" json:"name"`
	Photo  string         `form:"photo" json:"photo"`
	Detail string         `form:"detail" json:"detail"`
	Skus   []*entity.Skus `form:"skus" json:"skus"`
	Common
}

type ItemUpdate struct {
	ItemId string         `form:"item_id" json:"item_id"`
	Name   string         `form:"name" json:"name"`
	Photo  string         `form:"photo" json:"photo"`
	Detail string         `form:"detail" json:"detail"`
	Skus   []*entity.Skus `form:"skus" json:"skus"`
	Common
}

type ItemGet struct {
	ItemId string `form:"item_id" json:"item_id"`
	Fields string `form:"fields" json:"fields"`
	Common
}

type ItemDelete struct {
	ItemId        string `form:"item_id" json:"item_id"`
	IsFinalDelete *int   `form:"is_final_delete" json:"is_final_delete"`
	Common
}

type ItemRecover struct {
	ItemId string `form:"item_id" json:"item_id"`
	Common
}

type ItemSync struct {
	ItemId   string `form:"item_id" json:"item_id"`
	SyncType string `form:"sync_type" json:"sync_type"`
	Common
}

type ItemSearch struct {
	Fields    string `form:"fields" json:"fields"`
	ItemId    string `form:"item_id" json:"item_id"`
	ItemName  string `form:"item_name" json:"item_name"`
	SkuName   string `form:"sku_name" json:"sku_name"`
	Barcode   string `form:"barcode" json:"barcode"`
	ItemState *int   `form:"item_state" json:"item_state"`
	Page      int    `form:"page" json:"page"`
	Limit     int    `form:"limit" json:"limit"`
	Common
}
