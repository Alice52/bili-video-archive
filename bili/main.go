package main

import (
	"github.com/alice52/archive/common/global"
	initialize "github.com/alice52/archive/common/init"
	"github.com/alice52/archive/common/migration"
	"github.com/wordpress-plus/kit-logger/viperx"
	"github.com/wordpress-plus/kit-logger/zapx"
)

func main() {
	// init viper
	global.VIPER = viperx.Viper(&global.CONFIG) // 初始化Viper
	global.LOG = zapx.Zap(global.CONFIG.Zap)

	// init db and do migration
	global.DB = initialize.GormPgSQL()
	migration.Initialize(global.DB)
}
