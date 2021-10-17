package system

import "self-discipline/global"

type SysNickname struct {
	global.MODEL
	Nickname string `json:"nickname"` // 用户昵称
	Status   uint8  `json:"-"`        // 状态 0正常 1封号
}
