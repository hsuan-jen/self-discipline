package article

import (
	"self-discipline/global"
	"self-discipline/model/article"
	articleReq "self-discipline/model/article/request"

	"gorm.io/gorm"
)

func (h *Handler) Leave(m *article.ArticleLeaves) error {

	err := global.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&article.Articles{}).Where("id = ?", m.ArticleID).UpdateColumn("leave", gorm.Expr("`leave` + ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.Create(&m).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

func (h *Handler) LeaveList(m *articleReq.LeaveList) ([]article.ArticleLeaves, error) {

	var leaves []article.ArticleLeaves
	limit := m.PageSize
	offset := m.PageSize * (m.Page - 1)

	err := global.DB.Where("article_id = ?", m.ArticleID).
		Limit(limit).
		Offset(offset).
		Find(&leaves).Error

	return leaves, err
}
