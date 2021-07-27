package global

import (
	"self-discipline/pkg"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG pkg.Server
	VP     *viper.Viper
	LOG    *zap.Logger
)
