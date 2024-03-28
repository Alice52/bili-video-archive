package api

import (
	"fmt"
	c "github.com/alice52/archive/bili/api/common"
	"github.com/alice52/archive/bili/util"
	"github.com/alice52/archive/common/global"
	"github.com/wordpress-plus/kit-logger/viperx"
	"github.com/wordpress-plus/kit-logger/zapx"
	"testing"
)

func init() {
	global.VIPER = viperx.Viper(&global.CONFIG, "../config-local.yaml")
	global.LOG = zapx.Zap(global.CONFIG.Zap)
}

func TestPlayUrl(t *testing.T) {
	info, err := logonFunc().PlayUrl("BV1CA411S7q4", 1031064040, c.Qn4k, c.FnvalHDR|c.Fnval4K)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(util.MustMarshal(info))
}
