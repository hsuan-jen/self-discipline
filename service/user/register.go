package user

import (
	"errors"
	"self-discipline/global"
	"self-discipline/model/userInfo"
	"self-discipline/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (h *BaseService) Register(u *userInfo.Users) (err error, userInter *userInfo.Users) {

	var user userInfo.Users

	if !errors.Is(global.DB.Select("phone").Where("phone = ?", u.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		// 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}

	//附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V(u.Password)
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return err, u
}
