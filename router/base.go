package router

import (
	v1 "self-discipline/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (h *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	var baseApi = v1.ApiGroupApp.UserApiGroup.BaseApi
	BaseRouter := Router.Group("v1")
	{
		BaseRouter.POST("login", baseApi.Login)
		BaseRouter.POST("register", baseApi.Register)
	}
	return BaseRouter
}
