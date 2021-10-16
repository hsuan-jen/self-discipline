package service

import (
	"self-discipline/service/article"
	"self-discipline/service/system"
	"self-discipline/service/user"
)

type ServiceGroup struct {
	ArticleServiceGroup article.ServiceGroup
	UserServiceGroup    user.ServiceGroup
	SystemServiceGroup  system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
