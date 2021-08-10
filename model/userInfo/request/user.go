package request

// User login structure
type Login struct {
	Phone    string `json:"phone" form:"phone"`       // 手机号码
	Password string `json:"password" form:"password"` // 密码
}

type Register struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

type Avatar struct {
	Avatar string `json:"avatar" form:"avatar"`
}

type WechatLogin struct {
	Code string `json:"code" form:"code"`
}
