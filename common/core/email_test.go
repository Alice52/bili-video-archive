package core

import (
	"github.com/alice52/archive/common/global"
	"github.com/alice52/archive/common/util"
	"testing"
)

func TestEmail(_ *testing.T) {
	global.VIPER = Viper()

	err := util.EmailTest("Test Email", "Test Body")
	if err != nil {
		return
	}
}
