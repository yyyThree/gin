package service

import (
	"gin/dao"
	"gin/helper"
	"gin/library/valid"
	"gin/model/entity"
	"gin/model/param"
	"gin/output"
	"gin/output/code"
)

type Item struct {
}

func (item *Item) Add(params param.ItemAdd) (data *entity.Items, err error){
	if err = item.checkAdd(params); err != nil {
		return
	}

	// step1 生成商品ID
	itemId := helper.GenerateUuid()

	// step2 创建商品
	itemEntity := &entity.Items{
		Appkey: params.AppKey,
		Channel: params.Channel,
		ItemID: itemId,
		Name: params.Name,
		Photo: params.Photo,
		Detail: params.Detail,
	}
	data, err = dao.NewItem().Insert(itemEntity)

	if err != nil {
		return
	}

	return
}

// 参数校验
func (item *Item) checkAdd(params param.ItemAdd) error {
	myValid := valid.New()
	myValid.Append(valid.One("name", params.Name).NotEmpty().MaxStringLength(255))
	myValid.Append(valid.One("photo", params.Photo).NotEmpty().IsURL().MaxStringLength(512))

	if err := myValid.Valid(); err != "" {
		return output.Error(code.IllegalParams).WithDetails(err)
	}
	return nil
}