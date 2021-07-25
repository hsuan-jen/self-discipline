package router

import (
	"self-discipline/api/v1/user_handler"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	userHandler := user_handler.New()
	BaseRouter := Router.Group("v1")
	{
		BaseRouter.POST("login", userHandler.Login)
		BaseRouter.POST("register", userHandler.Register)

	}
	return BaseRouter
}
