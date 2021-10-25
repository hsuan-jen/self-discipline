package target

import (
	"self-discipline/global"
)

type Target struct {
	global.MODEL
	UserID    uint64 `json:"user_id"`    //用户id
	Title     string `json:"title"`      //目标标题
	Motive    string `json:"motive"`     //动机
	Plan      string `json:"plan"`       //计划
	Cycle     uint32 `json:"cycle"`      //计划周期(天数)
	Rest      uint32 `json:"rest"`       //休息天数
	StartTime int64  `json:"start_time"` //首次签到时间
	Amount    uint64 `json:"amount"`     //挑战金额
	Schedule  uint32 `json:"schedule"`   //计划进度
	Onlookers uint32 `json:"onlookers"`  //围观
	Care      uint32 `json:"care"`       //亲密度
	Sort      uint32 `json:"-"`          //排序 顺序
}
