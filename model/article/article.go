package article

import (
	"self-discipline/global"
	"self-discipline/model/userInfo"
)

type Articles struct {
	global.MODEL
	UserID    uint64         `json:"user_id"`
	Give      int32          `json:"give"`
	Like      int32          `json:"like"`
	Leave     int32          `json:"leave"`
	Content   string         `json:"content"`
	User      userInfo.Users `json:"user"`
	GiveMark  ArticleGives   `json:"give_mark" gorm:"foreignKey:ArticleID"`
	LeaveMark ArticleLeaves  `json:"leave_mark" gorm:"foreignKey:ArticleID"`
}

type ArticleGives struct {
	global.MODEL
	UserID    uint64 `json:"user_id"`
	ArticleID uint64 `json:"article_id"`
}

// 留言
type ArticleLeaves struct {
	global.MODEL
	UserID    uint64 `json:"user_id"`
	ArticleID uint64 `json:"article_id"`
	Msg       string `json:"msg"`
	PID       uint64 `json:"pid"`
}
