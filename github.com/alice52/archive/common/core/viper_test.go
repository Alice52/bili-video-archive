package core

import (
	"fmt"
	"github.com/alice52/archive/common/global"
	"testing"
)

func TestViper(_ *testing.T) {
	Viper()

	fmt.Printf("%#v", global.CONFIG)
}
