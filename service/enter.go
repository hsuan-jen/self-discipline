package service

import (
	"self-discipline/service/article"
	"self-discipline/service/user"
)

type ServiceGroup struct {
	ArticleServiceGroup article.ServiceGroup
	UserServiceGroup    user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
