package article

import (
	"self-discipline/global"
	"self-discipline/model/article"
	articleReq "self-discipline/model/article/request"
)

func (h *Handler) GetList(m *articleReq.GetList) ([]article.Articles, error) {

	var articles []article.Articles
	limit := m.PageSize
	offset := m.PageSize * (m.Page - 1)

	err := global.DB.Limit(limit).
		Offset(offset).
		Preload("User").
		Order("created_at desc").
		Find(&articles).Error

	return articles, err
}
