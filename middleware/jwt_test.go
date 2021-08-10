package middleware

import (
	"self-discipline/global"
	"self-discipline/model/userInfo/request"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestParseToken(t *testing.T) {
	j := JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		//UUID:     "3c8e1a99-9a55-428a-9549-42bdb26a357b",
		ID:       1,
		NickName: "可乐加冰",
		Phone:    "17299998888",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "Commander",                                       // 签名的发行者
		},
	}
	token, _ := j.CreateToken(claims)
	t.Log(claims)
	t.Log(token)
	data, _ := j.ParseToken(token)

	t.Log(data)
}
