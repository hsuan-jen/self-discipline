package user_service

import (
	"fmt"
	"self-discipline/configs"
	"self-discipline/global"
	"self-discipline/model"
	"self-discipline/utils"
	"strconv"
	"time"
)

func (s *service) SaveUserRedis(u *model.Users) error {
	m, err := utils.Struct2Map(u)
	if err != nil {
		return err
	}
	rkey := configs.RedisKeyUser + strconv.FormatUint(u.ID, 10)
	m["uuid"] = fmt.Sprintf("%v", m["uuid"])
	err = global.REDIS.HMSet(rkey, m).Err()
	global.REDIS.Expire(rkey, 86400*time.Second)
	return err
}
