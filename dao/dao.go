package dao

import (
	"fmt"
	"gin/constant"
	"gin/model/db"
	"gorm.io/gorm"
	"strings"
)

type dao struct {
	Tx *gorm.DB
	DbName string
}

// 获取数据库连接（默认master）
func (dao *dao) GetDb() (*gorm.DB, error) {
	if dao.Tx != nil {
		return dao.Tx, nil
	}
	return db.GetMasterDB(dao.DbName)
}

// 获取数据库连接（slave）
func (dao *dao) GetSlaveDb() (*gorm.DB, error) {
	return db.GetSlaveDB(dao.DbName)
}

func GetTx(txs ...*gorm.DB) *gorm.DB {
	var tx *gorm.DB
	if len(txs) != 0 {
		tx = txs[0]
	} else {
		tx = nil
	}
	return tx
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
