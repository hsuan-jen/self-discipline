package request

// User login structure
type Login struct {
	Phone    string `json:"phone" form:"phone" validate:"required" label:"手机号码"`           // 手机号码
	Password string `json:"password" form:"password" validate:"required,min=6" label:"密码"` // 密码
}

type Register struct {
	Phone    string `json:"phone" form:"phone" validate:"required" label:"手机号码"`
	Password string `json:"password" form:"password" validate:"required,min=6" label:"密码"`
}

type Avatar struct {
	Avatar string `json:"avatar" form:"avatar" validate:"required" label:"头像"`
}

type WechatLogin struct {
	Code string `json:"code" form:"code"`
}
