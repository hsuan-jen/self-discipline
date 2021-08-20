package request

import "self-discipline/model/common/request"

type Issue struct {
	Content string `json:"content" form:"content" validate:"required" label:"内容"`
}

type GetList struct {
	Keyword string `json:"keyword" form:"keyword"`
	request.PageInfo
}

type Give struct {
	ArticleID uint64 `json:"article_id" form:"article_id" validate:"required" label:"目标ID"`
}

type Leave struct {
	ArticleID uint64 `json:"article_id" form:"article_id" validate:"required" label:"目标ID"`
	Msg       string `json:"msg" form:"msg" validate:"required" label:"留言"`
}

type LeaveList struct {
	ArticleID uint64 `json:"article_id" form:"article_id" validate:"required" label:"目标ID"`
	request.PageInfo
}
