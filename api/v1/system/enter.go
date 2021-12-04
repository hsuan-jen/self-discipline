package system

import "self-discipline/service"

type ApiGroup struct {
	FileApi
}

var fileService = service.ServiceGroupApp.SystemServiceGroup.FileService
