package config

type System struct {
	Env       string `mapstructure:"env" json:"env" yaml:"env"`                      // 环境值
	JasyptPwd string `mapstructure:"jasypt-pwd" json:"jasypt-pwd" yaml:"jasypt-pwd"` // 配置加解密值
	Addr      int    `mapstructure:"addr" json:"addr" yaml:"addr"`                   // 端口值

	// DbType 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`

	UseRedis bool `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"` // 使用redis
	UseMongo bool `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"` // 使用redis
}
