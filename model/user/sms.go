package user

import "self-discipline/global"

//手机短信
type PhoneSms struct {
	global.MODEL
	Phone  string `json:"phone"` //手机号码
	Code   string `json:"code"`  // 验证码
	Status uint8  `json:"-"`     // 状态
}
