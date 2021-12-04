package system

import (
	"mime/multipart"
	"strings"

	"self-discipline/model/system"
	"self-discipline/utils/upload"
)

type FileService struct {
}

//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader
//@return: err error, file model.ExaFileUploadAndDownload
func (e *FileService) UploadFile(header *multipart.FileHeader) (file system.SysFile, err error) {
	oss := upload.NewOss()
	filePath, key, err := oss.UploadFile(header)
	if err != nil {
		return
	}
	s := strings.Split(header.Filename, ".")
	file = system.SysFile{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	return
}
