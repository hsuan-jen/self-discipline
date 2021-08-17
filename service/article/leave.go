package article

import (
	"self-discipline/global"
	"self-discipline/model/article"

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
