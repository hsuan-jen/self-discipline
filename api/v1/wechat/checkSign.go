package wechat

import (
	"self-discipline/global"
	userInfoReq "self-discipline/model/userInfo/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData request.CheckSignature true "微信加密签名, 时间戳, 随机数, 随机字符串"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /checkSign[get]
func CheckSign(c *gin.Context) {

	var req userInfoReq.CheckSignature
	_ = c.ShouldBind(&req)
	global.LOG.Info("CheckSign",
		zap.Any("signature", req.Signature),
		zap.Any("timestamp", req.Timestamp),
		zap.Any("nonce", req.Nonce),
		zap.Any("echostr", req.Echostr),
	)
	c.String(200, req.Echostr)
}
