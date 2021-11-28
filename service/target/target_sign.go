package target

import (
	"errors"
	"self-discipline/global"
	"self-discipline/model/target"
	targetReq "self-discipline/model/target/request"

	"gorm.io/gorm"
)

type TargetSignService struct{}

//打卡数据
func (s *TargetSignService) GetTargetSignList(factor targetReq.TargetSign, userID uint) (list []target.TargetSign, total int64, err error) {
	limit := factor.PageSize
	offset := factor.PageSize * (factor.Page - 1)

	db := global.DB.Model(&target.TargetSign{}).Where("status = 0 and user_id = ?", userID)
	if factor.Type == 0 {
		db = db.Order("day desc, like desc")
	} else {
		db = db.Order("created_at desc")
	}

	if err = db.Count(&total).Error; err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id,nickname,avatar,gender")
		}).
		Find(&list).Error

	return
}

//获取发布目标数据
func (s *TargetSignService) FindTargetSignByID(id uint64, userID uint) (info target.TargetSign, err error) {

	err = global.DB.Model(&target.Target{}).Where("id = ? and user_id = ?", id, userID).Order("created_id desc").Find(&info).Error
	return
}

//发布打卡数据
func (s *TargetSignService) CreateTargetSign(factor target.TargetSignSimple, userID uint) (err error) {
	var targetService TargetService
	//检查目标是否存折
	if _, err = targetService.FindTargetByID(factor.TargetID, userID); err != nil {
		return
	}

	lastTargetSignInfo, err := s.FindTargetSignByID(factor.TargetID, userID)
	factor.Day = lastTargetSignInfo.Day + 1
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		factor.Day = 0
	}

	err = global.DB.Create(&factor).Error
	return
}
