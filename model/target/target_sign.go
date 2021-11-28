package target

import (
	"self-discipline/global"
	"self-discipline/model/user"
)

type TargetSign struct {
	global.MODEL
	UserID   uint64           `json:"user_id"`   //
	TargetID uint64           `json:"target_id"` //目标id
	Day      uint32           `json:"day"`       //打卡天数
	Content  string           `json:"content"`   //内容
	Like     uint32           `json:"like"`      //点赞
	Msg      string           `json:"msg"`       //留言
	Images   string           `json:"images"`    //图片集，最多9张
	Video    string           `json:"video"`     //视频，单个
	Status   uint8            `json:"-"`         //状态 0正常 1引起不适 2内容与话题无关 3投诉3.1辱骂攻击 3.2引战 3.3垃圾广告3.4色情低俗 3.5政治相关 3.6违法违规3.7泄露他人隐私
	Sort     uint32           `json:"-"`         //排序 顺序 N人举报后未审核即可降权
	User     user.UserSurface `json:"user" `     //用户信息
}

type TargetSignSimple struct {
	global.MODEL
	UserID   uint64 `json:"user_id"`                                    //
	TargetID uint64 `json:"target_id" validate:"required" label:"目标id"` //目标id
	Day      uint32 `json:"day"`                                        //打卡天数
	Content  string `json:"content" validate:"required" label:"内容"`     //内容
	Like     uint32 `json:"like"`                                       //点赞
	Msg      string `json:"msg"`                                        //留言
	Images   string `json:"images"`                                     //图片集，最多9张
	Video    string `json:"video"`                                      //视频，单个
}

func (TargetSignSimple) TableName() string {
	return "target_sign"
}
