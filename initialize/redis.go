package initialize

import (
	"self-discipline/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.LOG.Fatal("redis connect ping failed, err:", zap.Any("err", err))
		return nil
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
}
