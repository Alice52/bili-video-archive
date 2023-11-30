package core

import (
	"github.com/alice52/archive/common/global"
	"testing"
)

func TestZap(_ *testing.T) {
	global.VIPER = Viper()

	global.LOG = Zap()
	global.LOG.Info("asasas")
}
