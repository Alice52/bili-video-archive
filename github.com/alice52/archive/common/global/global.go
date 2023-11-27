package global

import (
	"github.com/alice52/archive/common/config"
	"github.com/alice52/archive/common/util"
	"sync"

	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GLOBAL_DB    *gorm.DB
	GLOBAL_MONGO *qmgo.QmgoClient

	GLOBAL_DBList map[string]*gorm.DB
	GLOBAL_CONFIG config.Server
	GLOBAL_VP     *viper.Viper
	GLOBAL_LOG    *zap.Logger
	GLOBAL_Timer  util.Timer = util.NewTimerTask()

	lock sync.RWMutex
)
