package initialize

import (
	_ "self-discipline/docs"
	"self-discipline/global"
	"self-discipline/middleware"
	"self-discipline/router"
	"self-discipline/utils/env"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Mode)
	//var Router = gin.Default()
	var Router = gin.New()
	Router.Use(middleware.GinRecovery(true))
	Router.Use(middleware.GinLogger())

	// register pprof to gin
	if global.CONFIG.System.Pprof {
		if !env.Active().IsPro() {
			pprof.Register(Router)
			global.LOG.Info("register pprof handler")
		}
	}

	//register prometheus
	if global.CONFIG.System.Promhttp {
		Router.GET("/metrics", gin.WrapH(promhttp.Handler()))
		global.LOG.Info("register promhttp handler")
	}

	if !env.Active().IsPro() {
		Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		global.LOG.Info("register swagger handler")
	}

	// register rate
	if global.CONFIG.System.Rate {
		Router.Use(middleware.Rate())
		global.LOG.Info("use middleware rate")
	}

	//Router.StaticFS(global.CONFIG.Local.Path, http.Dir(global.ONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.LOG.Info("use middleware cors")

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		//systemRouter.InitUserRouter(PrivateGroup)                   // 用户操作

	}
	global.LOG.Info("router register success")
	return Router
}
