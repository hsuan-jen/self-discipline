package user

import (
	"self-discipline/global"
	"self-discipline/model"
	"self-discipline/model/request"
	"self-discipline/model/response"
	userService "self-discipline/service/user"
	"self-discipline/utils"

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
func Register(c *gin.Context) {
	var req request.Register
	_ = c.ShouldBind(&req)

	if err := utils.Verify(req, utils.RegisterVerify); err != nil {
		response.FailWithVerify(err.Error(), c)
		return
	}

	u := model.Users{Phone: req.Phone, Password: req.Password}
	err, userReturn := userService.Register(&u)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Any("err", err))
		response.FailWithDetailed(response.UserResponse{User: userReturn}, response.UserCreateError, c)
		return
	}
	response.OkWithDetailed(userReturn, "注册成功", c)
}
