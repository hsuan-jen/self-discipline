package user

import (
	"self-discipline/global"
	"self-discipline/model/common/response"
	"self-discipline/model/userInfo"
	userInfoReq "self-discipline/model/userInfo/request"
	userInfoRes "self-discipline/model/userInfo/response"
	"self-discipline/utils/validator"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户注册
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData request.Login true "手机号码, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /v1/register [post]
func (h *BaseApi) Register(c *gin.Context) {
	var req userInfoReq.Register

	if ok, msg := validator.Verify(c, &req); !ok {
		response.FailWithMessage(msg, c)
		return
	}

	u := userInfo.Users{Phone: req.Phone, Password: req.Password}
	err, userReturn := userService.Register(&u)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Any("err", err))
		response.FailWithDetailed(userInfoRes.UserResponse{User: userReturn}, "创建用户失败", c)
		return
	}
	response.OkWithDetailed(userReturn, "注册成功", c)
}
