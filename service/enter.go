package service

import (
	"self-discipline/service/article"
	"self-discipline/service/gateway"
	"self-discipline/service/system"
	"self-discipline/service/target"
)

type ServiceGroup struct {
	ArticleServiceGroup article.ServiceGroup
	SystemServiceGroup  system.ServiceGroup
	GatewayServiceGroup gateway.ServiceGroup
	TargetServiceGroup  target.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
