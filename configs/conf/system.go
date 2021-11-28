package conf

type System struct {
	Mode          string `mapstructure:"mode" json:"mode" yaml:"mode"`                              // sets gin mode
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                              // 端口值
	Pprof         bool   `mapstructure:"pprof" json:"pprof" yaml:"pprof"`                           // pprof
	Promhttp      bool   `mapstructure:"promhttp" json:"promhttp" yaml:"promhttp"`                  // promhttp
	Rate          bool   `mapstructure:"rate" json:"rate" yaml:"rate"`                              // 限流开关
	MaxBurstSize  int    `mapstructure:"max-burstSize" json:"maxBurstSize" yaml:"max-burstSize"`    // 限流最大数
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // Oss类型
}
