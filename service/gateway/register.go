package gateway

import (
	"errors"
	"self-discipline/global"
	"self-discipline/model/user"
	"self-discipline/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RegisterService struct{}

//手机注册
func (*RegisterService) RegisterByPhone(u user.Users) (info user.Users, errMsg string, err error) {

	err = global.DB.Select("phone").Where("phone = ?", u.Phone).First(&info).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 判断手机号是否注册
		return info, "手机号已注册", errors.New("手机号已注册")
	}

	//附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V(u.Password)
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return u, errMsg, err
}
