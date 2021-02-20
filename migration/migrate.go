package migration

import "sync"

// 约定 校验执行和实际执行 方法
type migration interface {
	needMigrate() bool
	migrate()
}

var migrations = []migration{
	new(addTableItems),
}

// 仅执行一次
var _migrate sync.Once

// 执行所有迁移
func Migrate()  {
	_migrate.Do(func() {
		for _, migration := range migrations {
			if !migration.needMigrate() {
				continue
			}
			migration.migrate()
		}
	})
}
