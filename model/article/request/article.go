package request

import "self-discipline/model/common/request"

type Issue struct {
	Content string `json:"content" form:"content"`
}

type GetList struct {
	Keyword string `json:"keyword" form:"keyword"`
	request.PageInfo
}
