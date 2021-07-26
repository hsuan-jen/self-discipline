package router

import (
	userHandler "self-discipline/api/v1/user"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	BaseRouter := Router.Group("v1")
	{
		BaseRouter.POST("login", userHandler.Login)
		BaseRouter.POST("register", userHandler.Register)
	}
	return BaseRouter
}
