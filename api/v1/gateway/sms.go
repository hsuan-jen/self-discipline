package gateway

import (
	"errors"
	"self-discipline/model/common/response"
	"self-discipline/model/user"
	userReq "self-discipline/model/user/request"
	"self-discipline/utils/validator"

	"github.com/gin-gonic/gin"
)

type SmsApi struct{}

// @Tags Base
// @Summary 获取短信
// @accept application/json
// @Produce application/json
// @Param data body userReq.PhoneSms true "手机号码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /phone/sms [post]
func (*SmsApi) Sms(ctx *gin.Context) {

	var req userReq.PhoneSms
	_ = ctx.ShouldBindJSON(&req)

	if ok, msg := validator.Verify(ctx, &req); !ok {
		response.FailWithMessage(msg, ctx)
		return
	}

	data := user.PhoneSms{Phone: req.Phone}
	if err := smsService.CreateSms(data); err != nil {
		if errors.Is(err, errors.New("1小时内限制10条短信")) {
			response.FailWithMessage("1小时内限制10条短信", ctx)
			return
		}
		response.FailWithMessage("获取短信失败", ctx)
		return
	}
	response.Ok(ctx)
}
