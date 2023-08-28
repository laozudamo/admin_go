package config

type ServerConfig struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Log   Log   `mapstructure:"log" json:"log" yaml:"log"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Jwt   Jwt   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
