package upload

import (
	"mime/multipart"

	"self-discipline/global"
)

//@interface_name: OSS
//@description: OSS接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

//@function: NewOss
//@description: OSS接口
//@description: OSS的实例化方法
//@return: OSS
func NewOss() OSS {
	switch global.CONFIG.System.OssType {
	case "qiniu":
		return &Qiniu{}
	case "local":
		fallthrough
	default:
		return &Local{}
	}
}
