package user

import (
	"fmt"
	"self-discipline/configs"
	"self-discipline/global"
	"self-discipline/model/userInfo"
	"self-discipline/utils"
	"strconv"
	"time"
)

func SaveUserRedis(u *userInfo.Users) error {
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
