package article

import (
	"self-discipline/global"
	"self-discipline/model/article"
	articleReq "self-discipline/model/article/request"

	"gorm.io/gorm"
)

func (h *Handler) GetList(m *articleReq.GetList, UserID uint64) ([]article.Articles, error) {

	var articles []article.Articles
	limit := m.PageSize
	offset := m.PageSize * (m.Page - 1)

	err := global.DB.Limit(limit).
		Offset(offset).
		Preload("User").
		Preload("GiveMark", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", UserID)
		}).
		Preload("LeaveMark", func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id = ?", UserID)
		}).
		Order("created_at desc").
		Find(&articles).Error

	return articles, err
}
