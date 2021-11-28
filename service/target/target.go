package target

import (
	"self-discipline/global"
	"self-discipline/model/target"
)

type TargetService struct{}

//获取发布目标数据
func (s *TargetService) FindTargetByID(id uint64, userID uint) (info target.Target, err error) {

	err = global.DB.Model(&target.Target{}).Where("id = ? and user_id = ?", id, userID).Find(&info).Error
	return
}
