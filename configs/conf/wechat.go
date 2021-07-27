package conf

type Wechat struct {
	AppID     string `mapstructure:"appid" json:"appid" yaml:"appid"`
	AppSecret string `mapstructure:"appsecret" json:"appsecret" yaml:"appsecret"`
}
