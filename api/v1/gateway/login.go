package gateway

import (
	"fmt"
	"self-discipline/configs"
	"self-discipline/global"
	"self-discipline/model/common/request"
	"self-discipline/model/common/response"
	"self-discipline/model/user"
	userReq "self-discipline/model/user/request"
	userRes "self-discipline/model/user/response"
	"self-discipline/service"
	"self-discipline/utils"
	"self-discipline/utils/validator"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginGroup struct{}

// @Tags Base
// @Summary 手机登录登录
// @Produce  application/json
// @Param data body userReq.LoginByPhone true "手机号码, 密码, 手机型号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /v1/phone/login [post]
func (g *LoginGroup) LoginByPhone(ctx *gin.Context) {
	var req userReq.LoginByPhone
	_ = ctx.ShouldBindJSON(&req)

	if ok, msg := validator.Verify(ctx, &req); !ok {
		response.FailWithMessage(msg, ctx)
		return
	}

	data := user.Users{Phone: req.Phone, Password: req.Password, MobileModel: req.MobileModel}
	if user, err := loginService.LoginByPhone(data); err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		response.FailWithMessage("用户名不存在或者密码错误", ctx)
	} else {
		g.tokenNext(ctx, user)
	}

}

// 登录以后签发jwt
func (g *LoginGroup) tokenNext(c *gin.Context, info user.Users) {
	j := utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:     info.UUID,
		ID:       info.ID,
		NickName: info.Nickname,
		Phone:    info.Phone,
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

	var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
	rkey := utils.MergeStr([]string{configs.RedisKeyJWT, strconv.FormatUint(info.ID, 10)})

	if err := jwtService.SetRedisJWT(token, rkey); err != nil {
		global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
		return
	}

	if err = g.UserSaveRedis(&info); err != nil {
		global.LOG.Error("记录登录信息失败!", zap.Any("err", err))
		response.FailWithMessage("记录登录信息失败", c)
		return
	}

	response.OkWithDetailed(userRes.LoginResponse{
		User:      info,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)

}

func (*LoginGroup) UserSaveRedis(info *user.Users) error {
	m, err := utils.Struct2Map(info)
	if err != nil {
		return err
	}
	rkey := utils.MergeStr([]string{configs.RedisKeyUserID, strconv.FormatUint(info.ID, 10)})
	m["uuid"] = fmt.Sprintf("%v", m["uuid"])
	err = global.REDIS.HMSet(rkey, m).Err()
	global.REDIS.Expire(rkey, 86400*time.Second)
	return err
}
