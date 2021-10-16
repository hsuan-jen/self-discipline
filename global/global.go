package global

import (
	"self-discipline/configs"

	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG configs.Server
	VP     *viper.Viper
	LOG    *zap.Logger

	Concurrency_Control = &singleflight.Group{}
)
