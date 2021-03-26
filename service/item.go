package service

import (
	"gin/constant"
	"gin/dao"
	"gin/helper"
	"gin/library/log"
	"gin/library/valid"
	"gin/model/entity"
	"gin/model/field"
	"gin/model/param"
	"gin/output"
	"gin/output/code"
	"strings"
)

type Item struct {
}

type ItemDetail struct {
	entity.Items
	Skus []*entity.Skus `json:"skus,omitempty" exTableName:"skus"`
}

func (item *Item) Add(params param.ItemAdd) (data ItemDetail, err error) {
	if err = item.checkAdd(params); err != nil {
		return
	}

	// step1 生成商品ID
	itemId := helper.GenerateUuid()

	// 开启事务
	db, err := dao.NewItem().GetDb()
	if err != nil {
		return
	}
	tx := db.Begin()
	defer tx.Commit()

	// step2 创建商品
	itemEntity := &entity.Items{
		Appkey:  params.AppKey,
		Channel: params.Channel,
		ItemID:  itemId,
		Name:    params.Name,
		Photo:   params.Photo,
		Detail:  params.Detail,
	}
	data.Items, err = dao.NewItem(tx).Insert(itemEntity)

	if err != nil {
		tx.Rollback()
		err = output.Error(code.ItemInsertFail)
		return
	}

	// step3 批量创建sku
	for _, sku := range params.Skus {
		sku.Appkey = params.AppKey
		sku.Channel = params.Channel
		sku.ItemID = itemId
		sku.SkuID = helper.GenerateUuid()
	}
	data.Skus, err = dao.NewSku(tx).InsertBatch(params.Skus)
	if err != nil {
		tx.Rollback()
		err = output.Error(code.SkuInsertFail)
		return
	}

	return
}

func (item *Item) Update(params param.ItemUpdate) (err error) {
	if err = item.checkUpdate(params); err != nil {
		return
	}

	// step1 获取商品信息
	itemInfo, err := item.Get(&param.ItemGet{
		ItemId: params.ItemId,
		Fields: "item_id,state,skus.sku_id",
		Common: params.Common,
	})
	if err != nil {
		return
	}

	// step2 校验商品是否可以更新
	if itemInfo.State != constant.ItemStateNormal {
		err = output.Error(code.ItemStateError)
		return
	}

	// 开启事务
	db, err := dao.NewItem().GetDb()
	if err != nil {
		return
	}
	tx := db.Begin()
	defer tx.Commit()

	// step3 更新商品信息
	err = dao.NewItem(tx).Update(constant.BaseMap{
		"name":   params.Name,
		"photo":  params.Photo,
		"detail": params.Detail,
	}, constant.BaseMap{
		"appkey":  params.AppKey,
		"channel": params.Channel,
		"item_id": params.ItemId,
	}, 1)
	if err != nil {
		tx.Rollback()
		err = output.Error(code.ItemUpdateFail)
		return
	}

	// step4 添加/更新/删除sku
	var skusNeedInsert []*entity.Skus
	var skuIdsNew []string
	skusNew := make(map[string]*entity.Skus)
	for _, sku := range params.Skus {
		sku.Appkey = params.AppKey
		sku.Channel = params.Channel
		sku.ItemID = params.ItemId
		if sku.SkuID == "" {
			sku.SkuID = helper.GenerateUuid()
			skusNeedInsert = append(skusNeedInsert, sku)
		} else {
			skuIdsNew = append(skuIdsNew, sku.SkuID)
			skusNew[sku.SkuID] = sku
		}
	}
	var skuIdsNow []string
	for _, sku := range itemInfo.Skus {
		skuIdsNow = append(skuIdsNow, sku.SkuID)
	}

	skuIdsNeedDel := helper.SliceDiff(skuIdsNow, skuIdsNew)
	skuIdsNeedUpdate := helper.SliceIntersect(skuIdsNow, skuIdsNew)

	// step4.1 添加sku
	if !helper.IsEmpty(skusNeedInsert) {
		_, err = dao.NewSku(tx).InsertBatch(skusNeedInsert)
		if err != nil {
			tx.Rollback()
			err = output.Error(code.SkuInsertFail)
			return
		}
	}

	// step4.2 更新sku
	if !helper.IsEmpty(skuIdsNeedUpdate) {
		for _, skuId := range skuIdsNeedUpdate {
			sku := skusNew[skuId.(string)]
			err = dao.NewSku(tx).Update(constant.BaseMap{
				"name":    sku.Name,
				"photo":   sku.Photo,
				"Barcode": sku.Barcode,
			}, constant.BaseMap{
				"appkey":  params.AppKey,
				"channel": params.Channel,
				"item_id": params.ItemId,
				"sku_id":  sku.SkuID,
			}, 1)
			if err != nil {
				tx.Rollback()
				err = output.Error(code.SkuUpdateFail)
				return
			}
		}
	}

	// step4.3 删除sku（业务上）
	if !helper.IsEmpty(skuIdsNeedDel) {
		err = dao.NewSku(tx).Update(constant.BaseMap{
			"state": constant.SkuStateDeletedSelf,
		}, constant.BaseMap{
			"appkey":  params.AppKey,
			"channel": params.Channel,
			"sku_id":  skuIdsNeedDel,
		}, len(skuIdsNeedDel))
		if err != nil {
			tx.Rollback()
			err = output.Error(code.SkuDelFail)
			return
		}
	}
	return
}

func (item *Item) Get(params *param.ItemGet) (data ItemDetail, err error) {
	if err = item.checkGet(params); err != nil {
		return
	}

	// step1 获取字段设置
	params.Fields = helper.GetString(params.Fields, "*")
	fields := helper.GetVerifyField(field.GetItemFields()["base"], params.Fields)
	if fields != "*" && !strings.Contains(params.Fields, "item_id") {
		params.Fields = strings.Join([]string{params.Fields, "item_id"}, ",")
	}

	// step2 按需读取数据
	baseWhere := constant.BaseMap{
		"appkey":  params.AppKey,
		"channel": params.Channel,
	}
	for _, exTableName := range field.ItemExTablesSort {
		getFields := strings.Split(helper.GetVerifyField(field.GetItemFields()[exTableName], params.Fields), ",")
		if getFields[0] == "" {
			continue
		}
		switch exTableName {
		case "base":
			item, _ := dao.NewItem().GetOne(getFields, helper.MergeMap(baseWhere, constant.BaseMap{
				"item_id": params.ItemId,
			}))
			if helper.IsEmpty(item) {
				err = output.Error(code.ItemNoFound)
				return
			}
			data.Items = item
		case "skus":
			skus, _ := dao.NewSku().GetList(getFields, helper.MergeMap(baseWhere, constant.BaseMap{
				"item_id": params.ItemId,
				"state":   []int{constant.SkuStateNormal, constant.SkuStateDeleted, constant.SkuStateFinalDeleted},
			}), constant.CommonLimit, 0)
			data.Skus = skus
		}
	}
	return
}

func (item *Item) Delete(params param.ItemDelete) (err error) {
	if err = item.checkDelete(params); err != nil {
		return
	}

	// step1 获取商品信息
	itemInfo, err := item.Get(&param.ItemGet{
		ItemId: params.ItemId,
		Fields: "item_id,state",
		Common: params.Common,
	})
	if err != nil {
		return
	}

	// step2 校验商品是否可以删除
	canDel := false
	state := constant.ItemStateDeleted
	switch *params.IsFinalDelete {
	case constant.ItemDelete: // 普通删除
		canDel = itemInfo.State == constant.ItemStateNormal
	case constant.ItemFinalDelete: // 彻底删除
		canDel = itemInfo.State == constant.ItemStateDeleted
		state = constant.ItemStateFinalDeleted
	}
	if !canDel {
		err = output.Error(code.ItemStateError)
		return
	}

	// step3 删除商品
	err = dao.NewItem().Update(constant.BaseMap{
		"state": state,
	}, constant.BaseMap{
		"appkey":  params.AppKey,
		"channel": params.Channel,
		"item_id": params.ItemId,
	}, 1)
	if err != nil {
		err = output.Error(code.ItemDelFail)
		log.New().Error("itemDelete", constant.BaseMap{
			"appkey":  params.AppKey,
			"channel": params.Channel,
			"item_id": params.ItemId,
			"err":     err,
		})
		return
	}
	log.New().Info("itemDelete", constant.BaseMap{
		"appkey":  params.AppKey,
		"channel": params.Channel,
		"item_id": params.ItemId,
	})

	// step4 删除sku
	go func() {
		_ = dao.NewSku().Update(constant.BaseMap{
			"state": state,
		}, constant.BaseMap{
			"appkey":  params.AppKey,
			"channel": params.Channel,
			"item_id": params.ItemId,
			"state":   []int{constant.SkuStateNormal, constant.SkuStateDeleted},
		}, constant.CommonLimit)
	}()

	return
}

func (item *Item) Recover(params param.ItemRecover) (err error) {
	if err = item.checkRecover(params); err != nil {
		return
	}

	// step1 获取商品信息
	itemInfo, err := item.Get(&param.ItemGet{
		ItemId: params.ItemId,
		Fields: "item_id,state",
		Common: params.Common,
	})
	if err != nil {
		return
	}

	// step2 校验商品是否可以恢复
	if !helper.InSlice(itemInfo.State, []int{constant.ItemStateDeleted, constant.ItemStateFinalDeleted}) {
		err = output.Error(code.ItemStateError)
		return
	}

	// step3 恢复商品
	err = dao.NewItem().Update(constant.BaseMap{
		"state": constant.ItemStateNormal,
	}, constant.BaseMap{
		"appkey":  params.AppKey,
		"channel": params.Channel,
		"item_id": params.ItemId,
	}, 1)
	if err != nil {
		err = output.Error(code.ItemDelFail)
		log.New().Error("itemRecover", constant.BaseMap{
			"appkey":  params.AppKey,
			"channel": params.Channel,
			"item_id": params.ItemId,
			"err":     err,
		})
		return
	}
	log.New().Info("itemRecover", constant.BaseMap{
		"appkey":  params.AppKey,
		"channel": params.Channel,
		"item_id": params.ItemId,
	})

	// step4 恢复sku
	go func() {
		_ = dao.NewSku().Update(constant.BaseMap{
			"state": constant.SkuStateNormal,
		}, constant.BaseMap{
			"appkey":  params.AppKey,
			"channel": params.Channel,
			"item_id": params.ItemId,
			"state":   []int{constant.SkuStateDeleted, constant.SkuStateFinalDeleted},
		}, constant.CommonLimit)
	}()

	return
}

// 参数校验
func (item *Item) checkAdd(params param.ItemAdd) error {
	myValid := valid.New()
	myValid.Append(valid.One("name", params.Name).NotEmpty().MaxStringLength(255))
	myValid.Append(valid.One("photo", params.Photo).NotEmpty().IsURL().MaxStringLength(512))
	myValid.Append(valid.One("skus", params.Name).NotEmpty())
	for _, sku := range params.Skus {
		myValid.Append(valid.One("skus.name", sku.Name).NotEmpty().MaxStringLength(255))
		myValid.Append(valid.One("skus.photo", sku.Photo).NotEmpty().IsURL().MaxStringLength(512))
		myValid.Append(valid.One("skus.barcode", sku.Barcode).NotEmpty())
	}

	if err := myValid.Valid(); err != "" {
		return output.Error(code.IllegalParams).WithDetails(err)
	}
	return nil
}

func (item *Item) checkUpdate(params param.ItemUpdate) error {
	myValid := valid.New()
	myValid.Append(valid.One("item_id", params.ItemId).NotEmpty().MaxStringLength(64))
	myValid.Append(valid.One("name", params.Name).NotEmpty().MaxStringLength(255))
	myValid.Append(valid.One("photo", params.Photo).NotEmpty().IsURL().MaxStringLength(512))
	myValid.Append(valid.One("skus", params.Name).NotEmpty())
	for _, sku := range params.Skus {
		myValid.Append(valid.One("skus.name", sku.Name).NotEmpty().MaxStringLength(255))
		myValid.Append(valid.One("skus.photo", sku.Photo).NotEmpty().IsURL().MaxStringLength(512))
		myValid.Append(valid.One("skus.barcode", sku.Barcode).NotEmpty())
	}

	if err := myValid.Valid(); err != "" {
		return output.Error(code.IllegalParams).WithDetails(err)
	}
	return nil
}

func (item *Item) checkGet(params *param.ItemGet) error {
	myValid := valid.New()
	myValid.Append(valid.One("item_id", params.ItemId).NotEmpty().MaxStringLength(64))

	if err := myValid.Valid(); err != "" {
		return output.Error(code.IllegalParams).WithDetails(err)
	}
	return nil
}

func (item *Item) checkDelete(params param.ItemDelete) error {
	myValid := valid.New()
	myValid.Append(valid.One("item_id", params.ItemId).NotEmpty().MaxStringLength(64))
	myValid.Append(valid.One("is_final_delete", params.IsFinalDelete).NotEmpty())
	if !helper.IsEmpty(params.IsFinalDelete) {
		myValid.Append(valid.One("is_final_delete", *params.IsFinalDelete).IsIn([]int{constant.ItemDelete, constant.ItemFinalDelete}))
	}

	if err := myValid.Valid(); err != "" {
		return output.Error(code.IllegalParams).WithDetails(err)
	}
	return nil
}

func (item *Item) checkRecover(params param.ItemRecover) error {
	myValid := valid.New()
	myValid.Append(valid.One("item_id", params.ItemId).NotEmpty().MaxStringLength(64))

	if err := myValid.Valid(); err != "" {
		return output.Error(code.IllegalParams).WithDetails(err)
	}
	return nil
}
