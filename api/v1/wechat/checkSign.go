package wechat

import (
	"github.com/gin-gonic/gin"
)

// @Tags Base
// @Summary 用户登录
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData request.CheckSignature true "微信加密签名, 时间戳, 随机数, 随机字符串"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /checkSign[get]
func (h *Oauth2Api) CheckSign(c *gin.Context) {

}
