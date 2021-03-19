package param

import "gin/model/entity"

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
	IsFinalDelete *int    `form:"is_final_delete" json:"is_final_delete"`
	Common
}

type ItemRecover struct {
	ItemId        string `form:"item_id" json:"item_id"`
	Common
}
