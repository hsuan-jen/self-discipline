package conf

type System struct {
	Addr     int  `mapstructure:"addr" json:"addr" yaml:"addr"`             // 端口值
	Pprof    bool `mapstructure:"pprof" json:"pprof" yaml:"pprof"`          // pprof
	Promhttp bool `mapstructure:"promhttp" json:"promhttp" yaml:"promhttp"` // promhttp
}
