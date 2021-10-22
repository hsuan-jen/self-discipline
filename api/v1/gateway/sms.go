package gateway

import (
	"self-discipline/global"
	"self-discipline/model/common/response"
	"self-discipline/model/user"
	userReq "self-discipline/model/user/request"
	"self-discipline/utils/validator"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SmsApi struct{}

// @Tags Base
// @Summary 获取短信
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData request.Sms true "手机号码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /v1/phone/sms [post]
func (*SmsApi) Sms(ctx *gin.Context) {

	var req userReq.PhoneSms
	_ = ctx.ShouldBindJSON(&req)

	if ok, msg := validator.Verify(ctx, &req); !ok {
		response.FailWithMessage(msg, ctx)
		return
	}

	data := user.PhoneSms{Phone: req.Phone}
	if errMsg, err := smsService.CreateSms(data); err != nil {
		global.LOG.Error("获取短信失败", zap.Any("err", err))
		if errMsg == "" {
			errMsg = "获取短信失败"
		}
		response.FailWithMessage(errMsg, ctx)
		return
	}
	response.Ok(ctx)
}
