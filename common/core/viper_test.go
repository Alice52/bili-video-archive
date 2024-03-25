package core

import (
	"fmt"
	"github.com/alice52/archive/common/global"
	"github.com/alice52/jasypt-go"
	"github.com/alice52/jasypt-go/constant"
	"os"
	"testing"
)

func TestViper(_ *testing.T) {
	Viper()

	_ = os.Setenv(constant.JasyptKey, constant.JasyptPwd)

	fmt.Printf("%#v", global.CONFIG)
}

func TestEncrypt(_ *testing.T) {
	etor := jasypt.New()

	wrapper, _ := etor.EncryptWrapper("frps.hubby.top")
	fmt.Printf("%#v", wrapper)
}
