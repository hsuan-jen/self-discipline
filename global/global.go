package global

import (
	"self-discipline/config"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VP     *viper.Viper
	LOG    *zap.Logger
	//GLOBAL_Timer timer.Timer = timer.NewTimerTask()
)
