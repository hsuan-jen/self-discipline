package request

//手机登录
type LoginByPhone struct {
	Phone       string `json:"phone" form:"phone" validate:"required" label:"手机号码"`           // 手机号码
	Password    string `json:"password" form:"password" validate:"required,min=6" label:"密码"` // 密码
	MobileModel string `json:"mobile_model" form:"mobile_model"`                              // 手机型号
}

//手机注册
type RegisterByPhone struct {
	Phone      string `json:"phone" form:"phone" validate:"required" label:"手机号码"`
	Password   string `json:"password" form:"password" validate:"required,min=6" label:"密码"`
	ConfirmPwd string `json:"confirm_pwd" form:"confirm_pwd" validate:"required,min=6,eqfield=Password" label:"确认密码"`
	Code       string `json:"code" form:"code" validate:"required,eq=6" label:"验证码"`
}

//获取短信
type PhoneSms struct {
	Phone string `json:"phone" form:"phone" validate:"required" label:"手机号码"`
}
