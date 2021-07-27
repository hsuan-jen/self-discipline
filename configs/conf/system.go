package conf

type System struct {
	Addr  int  `mapstructure:"addr" json:"addr" yaml:"addr"`    // 端口值
	Pprof bool `mapstructure:"pprof" json:"pprof" yaml:"pprof"` // 多点登录拦截
}
