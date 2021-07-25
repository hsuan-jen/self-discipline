package initialize

import (
	_ "self-discipline/docs"
	"self-discipline/global"
	"self-discipline/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.StaticFS(global.CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.LOG.Info("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	/* PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitApiRouter(PrivateGroup) // 注册功能api路由
		router.InitJwtRouter(PrivateGroup) // jwt相关路由

	} */
	global.LOG.Info("router register success")
	return Router
}
