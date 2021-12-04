package router

import (
	v1 "self-discipline/api/v1"

	"github.com/gin-gonic/gin"
)

type SystemRouter struct {
}

//系统路由
func (SystemRouter) InitSystemRouter(Router *gin.RouterGroup) {

	var fileApi = v1.ApiGroupApp.SystemApiGroup.FileApi
	fileRouter := Router.Group("file")
	{
		fileRouter.POST("upload", fileApi.UploadFile) //上传文件
	}
}
