package system

import (
	"self-discipline/global"
	"time"
)

//从redis取jwt
func (h *JwtService) GetRedisJWT(id string) (redisJWT string, err error) {
	redisJWT, err = global.REDIS.Get(id).Result()
	return redisJWT, err
}

func (h *JwtService) SetRedisJWT(jwt, id string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second
	err = global.REDIS.Set(id, jwt, timer).Err()
	return err
}

func (h *JwtService) DelRedisJWT(id string) (err error) {
	err = global.REDIS.Del(id).Err()
	return err
}
