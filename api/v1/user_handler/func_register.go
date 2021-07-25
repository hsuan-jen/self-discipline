package user_handler

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

func (h *handler) Register(c *gin.Context) {
	var req request.Register
	_ = c.ShouldBind(&req)

	if err := utils.Verify(req, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	u := model.Users{Phone: req.Phone, Password: req.Password}
	userService := userService.New()
	err, userReturn := userService.Register(&u)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Any("err", err))
		response.FailWithDetailed(response.UserResponse{User: userReturn}, err.Error(), c)
		return
	}
	response.OkWithDetailed(userReturn, "注册成功", c)
}
