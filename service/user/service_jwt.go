package user_service

import (
	"self-discipline/global"
	"time"
)

//从redis取jwt
func (s *service) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.REDIS.Get(userName).Result()
	return err, redisJWT
}

func (s *service) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second
	err = global.REDIS.Set(userName, jwt, timer).Err()
	return err
}
