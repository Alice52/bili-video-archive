package core

import (
	"github.com/alice52/archive/common/global"
	"github.com/alice52/archive/common/util"
	"github.com/wordpress-plus/kit-logger/viperx"
	"testing"
)

func TestEmail(_ *testing.T) {
	global.VIPER = viperx.Viper(&global.CONFIG, "config.yaml")

	err := util.EmailTest("Test Email", "Test Body")
	if err != nil {
		return
	}
}
