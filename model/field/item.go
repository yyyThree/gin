package field

import (
	"gin/constant"
	"gin/helper"
	"gin/model/entity"
	"reflect"
	"strings"
)

var (
	itemTables = constant.TableMap{
		"base": {
			new(entity.Items).TableName(),
			entity.ItemsColumns,
		},
		"skus": {
			new(entity.Skus).TableName(),
			entity.SkusColumns,
		},
	}
	// 逻辑层调用顺序
	ItemExTablesSort = []string{"base", "skus"}
	itemFields      = make(constant.FieldMap)
)

func GetItemFields() constant.FieldMap {
	// 仅初始化一次
	if !helper.IsEmpty(itemFields) {
		return itemFields
	}

	for exTableName, table := range itemTables {
		columns := reflect.ValueOf(table.Columns)
		if !columns.IsValid() {
			continue
		}
		columnsNum := columns.NumField()
		fields := make([]string, 0, columnsNum)
		for i := 0; i < columnsNum; i++ {
			column := columns.Field(i).String()
			if len(column) == 0 {
				continue
			}
			if exTableName != "base" {
				column = strings.Join([]string{exTableName, column}, ".")
			}
			fields = append(fields, column)
		}
		itemFields[exTableName] = fields
	}

	return itemFields
}
