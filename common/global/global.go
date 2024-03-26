package global

import (
	"github.com/alice52/archive/common/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	VIPER  *viper.Viper
	LOG    *zap.Logger
	DB     *gorm.DB
)

const (
	DbMysql = "mysql"
	DbPgsql = "pgsql"
)
