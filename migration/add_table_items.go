package migration

import (
	"gin/constant"
	"gin/entity"
	"gin/model"
)

type addTableItems struct {
}

func (addTableItems *addTableItems) needMigrate() bool {
	item := &entity.Item{}
	dblink := model.GetMasterDB(constant.DbServiceItems)
	return !dblink.Migrator().HasTable(item)
}

func (addTableItems *addTableItems) migrate() {
	item := &entity.Item{}
	dblink := model.GetMasterDB(constant.DbServiceItems)
	if err := dblink.Migrator().CreateTable(item); err != nil {
		panic("数据库迁移失败：addTableItems - " + err.Error())
	}
}
