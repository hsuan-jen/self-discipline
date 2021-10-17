package user

type UserInfo struct {
	UserID      uint64 `json:"user_id"`      //用户ID
	Coin        uint64 `json:"coin"`         // 金币
	Money       uint64 `json:"money"`        // 余额
	CardLog     int    `json:"card_log"`     // 打卡日志
	Focus       int    `json:"focus"`        // 关注
	Fan         int    `json:"fan"`          // 粉丝
	Praise      uint8  `json:"praise"`       // 获赞
	Vip         uint8  `json:"vip"`          // 会员等级
	VipValidity int32  `json:"vip_validity"` // 会员有效期，时间戳
	CreatedAt   int32  `json:"created_at"`
	UpdatedAt   int32  `json:"updated_at"`
}
