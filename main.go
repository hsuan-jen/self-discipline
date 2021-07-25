package main

import (
	"context"
	"net/http"
	"self-discipline/core"
	"self-discipline/global"
	"self-discipline/initialize"
	"self-discipline/utils/shutdown"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {

	global.VP = core.Viper()
	global.LOG = core.Zap()
	global.REDIS = initialize.Redis()

	dbRepo := initialize.Gorm()
	global.DB = dbRepo.Get()
	srv := core.RunWindowsServer()
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOG.Error(err.Error())
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				global.LOG.Error(err.Error())
			}
		},

		// 关闭 db
		func() {
			if global.DB != nil {
				if err := dbRepo.Close(); err != nil {
					global.LOG.Error(err.Error())
				}
			}
		},

		// 关闭 redis
		func() {
			if global.REDIS != nil {
				if err := global.REDIS.Close(); err != nil {
					global.LOG.Error(err.Error())
				}
			}
		},
	)
}
