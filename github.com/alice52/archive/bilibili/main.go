package main

import (
	"github.com/alice52/archive/common/core"
	"github.com/alice52/archive/common/global"
	initialize "github.com/alice52/archive/common/init"
	"github.com/alice52/archive/common/migration"
)

func main() {

	// init viper
	global.VIPER = core.Viper() // 初始化Viper

	// init zap
	global.LOG = core.Zap()

	// init db and do migration
	global.DB = initialize.GormPgSQL()
	migration.Initialize(global.DB)

}
