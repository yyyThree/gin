package model

import (
	"errors"
	"fmt"
	"gin/config"
	"gin/constant"
	"gin/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

// 仅加载一次配置
var _load sync.Once

var _dbLinks = map[string]map[string]*gorm.DB{
	"master": make(map[string]*gorm.DB),
	"slave":  make(map[string]*gorm.DB),
}

func Load() {
	_load.Do(func() {
		if err := open("master"); err != nil {
			panic(err)
		}
		if err := open("slave"); err != nil {
			panic(err)
		}
	})
}

func open(db string) error {
	dbConfig := config.Database{}
	switch db {
	case "master":
		dbConfig = config.Config.Db.Master
	case "slave":
		dbConfig = config.Config.Db.Slave
	default:
		return errors.New("数据库配置不存在")
	}

	if helper.IsEmpty(dbConfig.User) || helper.IsEmpty(dbConfig.Password) || helper.IsEmpty(dbConfig.Host) || helper.IsEmpty(dbConfig.Port) {
		return errors.New(db + "-数据库配置不正确")
	}

	for _, dbName := range constant.DataBases {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbName,
		)
		dblink, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second,   // 慢 SQL 阈值
					LogLevel:      logger.Info, // Log level
				},
			),
		})
		if err != nil {
			return errors.New(db + "-数据库连接失败：" + err.Error())
		}

		// 配置连接池
		sqlDB, err := dblink.DB()
		if err != nil {
			return errors.New(db + "-数据库连接池配置失败：" + err.Error())
		}
		// 最大连接数
		if dbConfig.MaxOpenConnections > 0 {
			sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConnections)
		}
		// 最大空闲连接数，即始终保持连接的数量
		if dbConfig.MaxIdleConnections > 0 {
			sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConnections)
		}
		_dbLinks[db][dbName] = dblink
	}

	return nil
}

func getDB(db string, dbName string) (*gorm.DB, error) {
	if dbLink, ok := _dbLinks[db][dbName]; ok {
		return dbLink, nil
	}
	return nil, errors.New("数据库连接失败")
}

func GetMasterDB(dbName string) *gorm.DB {
	dblink, _ := getDB("master", dbName)
	return dblink
}

func GetSlaveDB(dbName string) *gorm.DB {
	dblink, _ := getDB("slave", dbName)
	return dblink
}
