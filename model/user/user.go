package user

import (
	"self-discipline/global"

	uuid "github.com/satori/go.uuid"
)

type Users struct {
	global.MODEL
	UUID        uuid.UUID `json:"uuid"`     // 用户UUID
	Phone       string    `json:"phone"`    // 手机号码
	Password    string    `json:"-"`        // 用户登录密码
	Nickname    string    `json:"nickname"` // 用户昵称
	Avatar      string    `json:"avatar"`   // 用户头像
	Gender      uint8     `json:"gender"`   // 性别
	Status      uint8     `json:"-"`        // 状态 0正常 1封号
	PayPwd      string    `json:"-"`        // 支付密码
	MobileModel string    `json:"-"`        // 手机型号
}