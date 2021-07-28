package user

import (
	"encoding/json"
	httpURL "net/url"
	"self-discipline/configs"
	"self-discipline/global"
	"self-discipline/middleware"
	"self-discipline/model"
	"self-discipline/model/request"
	"self-discipline/model/response"
	userService "self-discipline/service/user"
	"self-discipline/utils"
	"self-discipline/utils/httpclient"
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
func Login(c *gin.Context) {
	var req request.Login
	_ = c.ShouldBind(&req)

	if err := utils.Verify(req, utils.LoginVerify); err != nil {
		response.FailWithVerify(err.Error(), c)
		return
	}

	u := model.Users{Phone: req.Phone, Password: req.Password}
	if err, user := userService.Login(&u); err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		response.FailWithCode(response.UserEmpty, c)
	} else {
		tokenNext(c, *user)
	}
}

// @Tags Base
// @Summary 微信登录
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData request.WechatLogin true "token"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /v1/wechatlogin [post]
func WechatLogin(c *gin.Context) {
	var req request.WechatLogin
	_ = c.ShouldBind(&req)

	if err := utils.Verify(req, utils.WechatLoginVerify); err != nil {
		response.FailWithVerify(err.Error(), c)
		return
	}
	//获取access_token url
	var accessTokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token"
	//appid=%s&secret=%s&code=%s&grant_type=authorization_code

	form := make(httpURL.Values)
	form["appid"] = []string{global.CONFIG.Wechat.AppID}
	form["secret"] = []string{global.CONFIG.Wechat.AppSecret}
	form["code"] = []string{req.Code}
	form["grant_type"] = []string{"authorization_code"}

	body, err := httpclient.Get(accessTokenUrl, form, httpclient.WithTTL(time.Second*3))
	if err != nil {
		global.LOG.Error("登陆失败! 获取access_token错误!", zap.Any("err", err))
		response.FailWithCode(response.WechatAccessTokenErr, c)
	}

	type accessTokenInfo struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Openid       string `json:"openid"`
		Scope        string `json:"scope"`
	}
	res := new(accessTokenInfo)
	err = json.Unmarshal(body, res)
	if err != nil {
		global.LOG.Error("登陆失败! 解析access_token返回值错误!", zap.Any("err", err))
		response.FailWithCode(response.WechatAccessTokenErr, c)
	}

	//通过access_token获取userinfo url
	var userInfoUrl = "https://api.weixin.qq.com/sns/userinfo"
	//?access_token=ACCESS_TOKEN&openid=OPENID
	form = make(httpURL.Values)
	form["access_token"] = []string{res.AccessToken}
	form["openid"] = []string{res.Openid}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.Users) {
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
		response.FailWithCode(response.JwtCreateError, c)
		return
	}

	rkey := configs.RedisKeyJWT + strconv.FormatUint(user.ID, 10)

	if err := userService.SetRedisJWT(token, rkey); err != nil {
		global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
		response.FailWithCode(response.UserSetStatusErr, c)
		return
	}

	if err = userService.SaveUserRedis(&user); err != nil {
		global.LOG.Error("记录登录信息失败!", zap.Any("err", err))
		response.FailWithCode(response.UserRecordErr, c)
		return
	}

	response.OkWithDetailed(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)

}
