package userInfo

import (
	"self-discipline/global"

	uuid "github.com/satori/go.uuid"
)

type Users struct {
	global.MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                                 // 用户UUID
	Phone    string    `json:"phone" gorm:"comment:手机号码"`                                                  // 手机号码
	Password string    `json:"-"  gorm:"comment:用户登录密码"`                                                   // 用户登录密码
	Nickname string    `json:"nickname" gorm:"comment:用户昵称"`                                               // 用户昵称
	Avatar   string    `json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	Gender   uint8     `json:"gender" gorm:"comment:性别"`                                                   // 性别
}
