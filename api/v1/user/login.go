package user

import (
	"self-discipline/configs"
	"self-discipline/global"
	"self-discipline/middleware"
	"self-discipline/model/common/request"
	"self-discipline/model/common/response"
	"self-discipline/model/userInfo"
	userInfoReq "self-discipline/model/userInfo/request"
	userInfoRes "self-discipline/model/userInfo/response"
	"self-discipline/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData request.Login true "手机号码, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /v1/login [post]
func (h *BaseApi) Login(c *gin.Context) {
	var req userInfoReq.Login
	_ = c.ShouldBind(&req)

	if err := utils.Verify(req, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	u := userInfo.Users{Phone: req.Phone, Password: req.Password}
	if err, user := userService.Login(&u); err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		tokenNext(c, *user)
	}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user userInfo.Users) {
	j := &middleware.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.Nickname,
		Phone:    user.Phone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "Commander",                                       // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	rkey := configs.RedisKeyJWT + strconv.FormatUint(user.ID, 10)

	if err := userService.SetRedisJWT(token, rkey); err != nil {
		global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
		return
	}

	if err = userService.SaveUserRedis(&user); err != nil {
		global.LOG.Error("记录登录信息失败!", zap.Any("err", err))
		response.FailWithMessage("记录登录信息失败", c)
		return
	}

	response.OkWithDetailed(userInfoRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)

}
