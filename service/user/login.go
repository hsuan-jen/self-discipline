package user

import (
	"self-discipline/global"
	"self-discipline/model"
	"self-discipline/utils"
)

func Login(u *model.Users) (err error, userInter *model.Users) {
	var user model.Users
	u.Password = utils.MD5V(u.Password)
	err = global.DB.Where("phone = ? AND password = ?", u.Phone, u.Password).First(&user).Error
	return err, &user
}
