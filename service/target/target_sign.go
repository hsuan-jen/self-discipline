package target

import (
	"self-discipline/global"
	"self-discipline/model/target"
	targetReq "self-discipline/model/target/request"

	"gorm.io/gorm"
)

type TargetSignService struct{}

func (s *TargetSignService) GetTargetSignList(factor targetReq.TargetSign) (list []target.TargetSign, err error) {
	limit := factor.PageSize
	offset := factor.PageSize * (factor.Page - 1)

	db := global.DB.Model(&target.TargetSign{}).Where("status = 0")
	if factor.Type == 0 {
		db = db.Order("day desc, like desc")
	} else {
		db = db.Order("created_at desc")
	}

	err = db.Limit(limit).Offset(offset).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id,phone,nickname,avatar,gender")
		}).
		Find(&list).Error

	return
}
