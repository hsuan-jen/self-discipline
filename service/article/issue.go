package article

import (
	"self-discipline/global"
	"self-discipline/model/article"
)

func (h *Handler) AddArticle(m *article.Articles) error {

	err := global.DB.Create(&m).Error
	return err
}
