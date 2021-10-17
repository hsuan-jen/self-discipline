package service

import (
	"self-discipline/service/article"
	"self-discipline/service/gateway"
	"self-discipline/service/system"
)

type ServiceGroup struct {
	ArticleServiceGroup article.ServiceGroup
	SystemServiceGroup  system.ServiceGroup
	GatewayServiceGroup gateway.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
