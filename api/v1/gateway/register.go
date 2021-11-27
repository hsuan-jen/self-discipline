package gateway

import (
	"errors"
	"self-discipline/global"
	"self-discipline/model/common/response"
	"self-discipline/model/user"
	userReq "self-discipline/model/user/request"
	"self-discipline/utils/validator"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RegisterApi struct{}

// @Tags Base
// @Summary 手机用户注册
// @accept application/json
// @Produce application/json
// @Param data body userReq.RegisterByPhone true "手机号码, 密码, 确认密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /phone/register [post]
func (*RegisterApi) RegisterByPhone(ctx *gin.Context) {

	var req userReq.RegisterByPhone
	_ = ctx.ShouldBindJSON(&req)

	if ok, msg := validator.Verify(ctx, &req); !ok {
		response.FailWithMessage(msg, ctx)
		return
	}

	//短信验证
	sms := user.PhoneSms{Phone: req.Phone}
	smsInfo, err := smsService.FindSms(sms)
	if err != nil {
		global.LOG.Error("验证短信失败", zap.Any("err", err))
		response.FailWithMessage("验证短信失败", ctx)
		return
	}
	if smsInfo.Status == 1 {
		response.FailWithMessage("短信已使用", ctx)
		return
	} else if smsInfo.CreatedAt+300 < int32(time.Now().Unix()) {
		response.FailWithMessage("短信有效期为5分钟", ctx)
		return
	}
	if err := smsService.UpdateSms(smsInfo.ID); err != nil {
		global.LOG.Error("更新短信失败", zap.Any("err", err))
		response.FailWithMessage("更新短信失败", ctx)
		return
	}

	sysnickname, err := nicknameService.GetNickname()
	if err != nil {
		global.LOG.Error("获取昵称失败", zap.Any("err", err))
		response.FailWithMessage("获取昵称失败", ctx)
		return
	}
	u := user.User{
		Phone:    req.Phone,
		Password: req.Password,
		Nickname: sysnickname.Nickname,
		UserInfo: user.UserInfo{CreatedAt: int32(time.Now().Unix())},
	}
	res, err := registerService.RegisterByPhone(u)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 判断手机号是否注册
			response.FailWithMessage("手机号已注册", ctx)
			return
		}
		response.FailWithMessage("注册失败", ctx)
		return
	}
	response.OkWithDetailed(res, "注册成功", ctx)
}
