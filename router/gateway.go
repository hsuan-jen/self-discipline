package router

import (
	v1 "self-discipline/api/v1"

	"github.com/gin-gonic/gin"
)

type GatwayRouter struct {
}

func (GatwayRouter) InitGatwayRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	var loginApi = v1.ApiGroupApp.GatewayApiGroup.LoginApi
	var registerApi = v1.ApiGroupApp.GatewayApiGroup.RegisterApi
	var smsApi = v1.ApiGroupApp.GatewayApiGroup.SmsApi
	gatewayRouter := Router.Group("v1")
	{
		gatewayRouter.POST("phone/login", loginApi.LoginByPhone)
		gatewayRouter.POST("phone/register", registerApi.RegisterByPhone)
		gatewayRouter.POST("phone/sms", smsApi.Sms)
	}
	return gatewayRouter
}
