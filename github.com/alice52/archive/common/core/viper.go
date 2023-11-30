package core

import (
	"flag"
	"fmt"
	"github.com/alice52/archive/common/global"
	"github.com/alice52/jasypt-go"
	"github.com/alice52/jasypt-go/constant"
	jasyptv "github.com/alice52/jasypt-go/viper"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// Viper 优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	v := viper.NewWithOptions()
	v.SetConfigFile(getConfigFile(path...))
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		unmarshalConfig(v)
	})

	unmarshalConfig(v)

	return v
}

func getConfigFile(path ...string) string {
	var config string

	// parse from code
	if len(path) != 0 {
		config = path[0]
	}

	// parse from cmd
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()

	// parse from env: gin.mode

	// pickup default value
	if config == "" {
		config = ConfigDefaultFile
	}

	fmt.Printf("using viper config: %s", config)

	return config
}

// unmarshalConfig
//  1. parse config
//  2. do decrypt by jasypt
func unmarshalConfig(v *viper.Viper) {
	jasyptPwd := global.CONFIG.System.JasyptPwd
	if len(jasyptPwd) == 0 {
		jasyptPwd = constant.JasyptPwd
	}

	_ = os.Setenv(constant.JasyptKey, jasyptPwd)
	if err := jasyptv.Unmarshal(v, jasypt.New(), &global.CONFIG); err != nil {
		fmt.Println(err)
	}
}
