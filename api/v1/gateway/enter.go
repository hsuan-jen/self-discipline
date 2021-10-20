package gateway

import "self-discipline/service"

type ApiGroup struct {
	LoginGroup
	RegisterGroup
	SmsGroup
}

var loginService = service.ServiceGroupApp.GatewayServiceGroup.LoginService
var registerService = service.ServiceGroupApp.GatewayServiceGroup.RegisterService
var smsService = service.ServiceGroupApp.GatewayServiceGroup.SmsService
var nicknameService = service.ServiceGroupApp.SystemServiceGroup.NicknameService
