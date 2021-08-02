package conf

type System struct {
	Mode         string `mapstructure:"mode" json:"mode" yaml:"mode"`                           // sets gin mode
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`                           // 端口值
	Pprof        bool   `mapstructure:"pprof" json:"pprof" yaml:"pprof"`                        // pprof
	Promhttp     bool   `mapstructure:"promhttp" json:"promhttp" yaml:"promhttp"`               // promhttp
	Rate         bool   `mapstructure:"rate" json:"rate" yaml:"rate"`                           // 限流开关
	MaxBurstSize int    `mapstructure:"max-burstSize" json:"maxBurstSize" yaml:"max-burstSize"` // 限流最大数
}
