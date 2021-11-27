package router

import (
	v1 "self-discipline/api/v1"

	"github.com/gin-gonic/gin"
)

type TargetRouter struct {
}

func (r *TargetRouter) InitTargetRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	var targetApi = v1.ApiGroupApp.TargetApiGroup
	targetRouter := Router.Group("target")
	{
		targetRouter.GET("getTargetSignList", targetApi.TargetSignApi.GetTargetSignList)

	}
	return targetRouter
}
