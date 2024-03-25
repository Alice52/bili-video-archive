package api

import (
	"fmt"
	"github.com/alice52/archive/bilibili/util"
	"github.com/alice52/archive/common/global"
	"github.com/wordpress-plus/kit-logger/viperx"
	"testing"
)

func TestClient_IsLogin(t *testing.T) {
	client := &BClient{}
	if client.isLogin() {
		fmt.Println("client has already login")
	} else {
		fmt.Println("client has not login")
	}
}

func TestClient_MySpaceInfo(t *testing.T) {
	global.VIPER = viperx.Viper(&global.CONFIG, "../../config.yaml")

	client, err := GetLogonClient()
	if err != nil {
		t.Error(err)
		return
	}

	info, err := client.MySpaceInfo()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(util.MustMarshal(info))
}
