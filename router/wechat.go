package router

import (
	wechatHandler "self-discipline/api/v1/wechat"

	"github.com/gin-gonic/gin"
)

func InitWechatRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	WechatRouter := Router.Group("")
	{
		WechatRouter.GET("checkSign", wechatHandler.CheckSign)
	}
	return WechatRouter
}
