package main

import (
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/common/global"
	initialize "github.com/alice52/archive/common/init"
	"github.com/wordpress-plus/kit-logger/viperx"
	"github.com/wordpress-plus/kit-logger/zapx"
)

func main() {
	// init viper
	global.VIPER = viperx.Viper(&global.CONFIG, "config-local.yaml") // 初始化Viper
	global.LOG = zapx.Zap(global.CONFIG.Zap)

	// init db and do migration
	global.DB = initialize.GormMysql()
	if global.DB.Error != nil {
		panic(global.DB.Error)
	} else {
		dal.SetDefault(global.DB)
	}
}
