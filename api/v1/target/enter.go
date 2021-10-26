package target

import "self-discipline/service"

type ApiGroup struct {
	TargetSignApi
}

var targetSignService = service.ServiceGroupApp.TargetServiceGroup.TargetSignService
