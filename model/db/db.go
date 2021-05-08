package db

import (
	"errors"
	"fmt"
	"github.com/rubenv/sql-migrate"
	"github.com/yyyThree/gin/config"
	"github.com/yyyThree/gin/constant"
	"github.com/yyyThree/gin/helper"
	"github.com/yyyThree/gin/output"
	"github.com/yyyThree/gin/output/code"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"sync"
	"time"
)

// 仅加载一次配置
var load sync.Once

var dbLinks = map[string]map[string]*gorm.DB{
	"master": make(map[string]*gorm.DB),
	"slave":  make(map[string]*gorm.DB),
}

func Load() {
	load.Do(func() {
		if err := open("master"); err != nil {
			panic(err)
		}
		if err := open("slave"); err != nil {
			panic(err)
		}
		dbMigrate()
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
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbName,
		)
		dbLoggerConfig := logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
		}
		if config.IsDev() {
			dbLoggerConfig.LogLevel = logger.Info
		}
		dblink, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名，启用该选项后，`Item` 表将是`item`
			},
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				dbLoggerConfig,
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
		// 连接可复用的最大时间
		if dbConfig.MaxConnectionIdleTime > 0 {
			sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(dbConfig.MaxIdleConnections))
		}
		dbLinks[db][dbName] = dblink
	}

	return nil
}

func getDB(db string, dbName string) (*gorm.DB, error) {
	if dbLink, ok := dbLinks[db][dbName]; ok {
		return dbLink, nil
	}
	return nil, output.Error(code.MySqlErr)
}

func GetMasterDB(dbName string) (*gorm.DB, error) {
	dblink, err := getDB("master", dbName)
	return dblink, err
}

func GetSlaveDB(dbName string) (*gorm.DB, error) {
	dblink, err := getDB("slave", dbName)
	return dblink, err
}

// 执行数据库迁移
func dbMigrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: "migration",
	}
	dblink, _ := GetMasterDB(constant.DbServiceItems)
	sqlDB, _ := dblink.DB()
	_, err := migrate.Exec(sqlDB, "mysql", migrations, migrate.Up)
	if err != nil {
		panic("sqlMigrate err " + err.Error())
	}
	return
}
