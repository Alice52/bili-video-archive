package config

import (
	"github.com/wordpress-plus/kit-logger/zapx/config"
)

type Server struct {
	Zap    config.Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	Email  Email      `mapstructure:"email" json:"email" yaml:"email"`
	System System     `mapstructure:"system" json:"system" yaml:"system"`

	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
}
