package dao

import (
	"fmt"
	"gin/constant"
	"gorm.io/gorm"
	"strings"
)

type dao struct {
	*gorm.DB
}

// 封装whereIn
func (dao *dao) WhereIn(sql *gorm.DB, whereIn constant.SqlWhereInMap) *gorm.DB {
	for k, v := range whereIn {
		sql.Where(k+" In ?", v)
	}
	return sql
}

// 封装between
func (dao *dao) Between(sql *gorm.DB, between constant.SqlBetweenInMap) *gorm.DB {
	for k, v := range between {
		sql.Where(k+" BETWEEN ? AND ?", v[0], v[1])
	}
	return sql
}

// 封装like
func (dao *dao) Like(sql *gorm.DB, like constant.BaseMap) *gorm.DB {
	for k, v := range like {
		sql.Where(k+" LIKE ?", "%%"+fmt.Sprintf("%v", v)+"%%")
	}
	return sql
}

// 封装orderBy
func (dao *dao) OrderBy(sql *gorm.DB, orderBy constant.SqlOrderByMap) *gorm.DB {
	var order []string
	for k, v := range orderBy {
		order = append(order, k+" "+v)
	}
	if len(order) > 0 {
		sql.Order(strings.Join(order, ","))
	}
	return sql
}

