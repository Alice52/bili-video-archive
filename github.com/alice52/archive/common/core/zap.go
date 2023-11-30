package core

import (
	"fmt"
	"github.com/alice52/archive/common/core/internal"
	"github.com/alice52/archive/common/global"
	"github.com/alice52/archive/common/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	zd := global.CONFIG.Zap.Director
	if ok, _ := util.PathExists(zd); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", zd)
		_ = os.Mkdir(zd, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	zap.ReplaceGlobals(logger)

	return logger
}
