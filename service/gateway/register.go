package gateway

import (
	"self-discipline/global"
	"self-discipline/model/user"
	"self-discipline/utils"

	uuid "github.com/satori/go.uuid"
)

type RegisterService struct{}

//手机注册
func (*RegisterService) RegisterByPhone(u user.User) (info user.User, err error) {

	err = global.DB.Select("phone").Where("phone = ?", u.Phone).First(&info).Error
	if err != nil {
		return
	}

	//附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V(u.Password)
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return u, err
}
