package user

import (
	"self-discipline/global"
	"self-discipline/model/userInfo"
	"self-discipline/utils"
)

func Login(u *userInfo.Users) (err error, userInter *userInfo.Users) {
	var user userInfo.Users
	u.Password = utils.MD5V(u.Password)
	err = global.DB.Where("phone = ? AND password = ?", u.Phone, u.Password).First(&user).Error
	return err, &user
}
