package system

import (
	"self-discipline/global"
	"self-discipline/model/common/response"
	"self-discipline/model/system"
	systemRes "self-discipline/model/system/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileApi struct {
}

// @Tags File
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /file/upload [post]
func (u *FileApi) UploadFile(c *gin.Context) {
	var file system.SysFile
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileService.UploadFile(header) // 文件上传后拿到文件路径
	if err != nil {
		global.LOG.Error("上传失败!", zap.Error(err))
		response.FailWithMessage("上传失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysFileResponse{File: file}, "上传成功", c)
}
