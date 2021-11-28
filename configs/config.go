package configs

import (
	"self-discipline/configs/conf"
)

type Server struct {
	JWT    conf.JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    conf.Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  conf.Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  conf.Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System conf.System `mapstructure:"system" json:"system" yaml:"system"`
	// oss
	Local conf.Local `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu conf.Qiniu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
}
