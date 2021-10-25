package gateway

import (
	"self-discipline/global"
	"self-discipline/model/user"
	"self-discipline/utils"
)

type LoginService struct{}

func (*LoginService) LoginByPhone(u user.User) (info user.User, err error) {
	u.Password = utils.MD5V(u.Password)
	err = global.DB.Where("phone = ? AND password = ?", u.Phone, u.Password).First(&info).Error
	return info, err
}
