package user

//手机短信
type PhoneSms struct {
	Phone     string `json:"phone"` //手机号码
	Code      string `json:"code"`  // 验证码
	Status    uint8  `json:"-"`     // 状态
	CreatedAt int32  `json:"-"`
	UpdatedAt int32  `json:"-"`
}