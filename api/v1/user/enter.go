package user

import (
	"self-discipline/service"
)

type BaseApi struct{}

type ApiGroup struct {
	BaseApi
}

var userService = service.ServiceGroupApp.UserServiceGroup
