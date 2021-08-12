package router

import (
	v1 "self-discipline/api/v1"

	"github.com/gin-gonic/gin"
)

func InitWechatRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	var wechatRouter = v1.ApiGroupApp.WechatApiGroup.Oauth2Api
	WechatRouter := Router.Group("")
	{
		WechatRouter.GET("checkSign", wechatRouter.CheckSign)
	}
	return WechatRouter
}
