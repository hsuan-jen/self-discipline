package article

import (
	"errors"
	"self-discipline/global"
	"self-discipline/model/article"

	"gorm.io/gorm"
)

func (h *Handler) Give(UserID, ArticleID uint64) error {

	var give article.ArticleGives
	var db = global.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := db.Where("deleted_at = 0 and user_id = ? and article_id = ?", UserID, ArticleID).First(&give).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}

	if give == (article.ArticleGives{}) {
		db.Model(&article.Articles{}).Where("id = ?", ArticleID).UpdateColumn("give", gorm.Expr("give + ?", 1))
		err = db.Create(&article.ArticleGives{UserID: UserID, ArticleID: ArticleID}).Error
	} else {
		db.Model(&article.Articles{}).Where("id = ?", ArticleID).UpdateColumn("give", gorm.Expr("give - ?", 1))
		err = db.Where("user_id = ? and article_id = ?", UserID, ArticleID).Delete(&give).Error
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
