package global

import (
	"github.com/alice52/archive/common/config"
	"sync"

	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	VIPER  *viper.Viper
	LOG    *zap.Logger

	DB     *gorm.DB
	DBList map[string]*gorm.DB
	MONGO  *qmgo.QmgoClient

	Timer *config.Timer

	lock sync.RWMutex
)
