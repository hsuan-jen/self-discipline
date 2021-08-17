package request

import "self-discipline/model/common/request"

type Issue struct {
	Content string `json:"content" form:"content"`
}

type GetList struct {
	Keyword string `json:"keyword" form:"keyword"`
	request.PageInfo
}

type Give struct {
	ArticleID uint64 `json:"article_id" form:"article_id"`
}

type Leave struct {
	ArticleID uint64 `json:"article_id" form:"article_id"`
	Msg       string `json:"msg" form:"msg"`
}
