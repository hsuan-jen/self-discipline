package v1

import (
	"self-discipline/api/v1/article"
	"self-discipline/api/v1/gateway"
	"self-discipline/api/v1/system"
	"self-discipline/api/v1/target"
	"self-discipline/api/v1/wechat"
)

type ApiGroup struct {
	WechatApiGroup  wechat.ApiGroup
	ArticleApiGroup article.ApiGroup
	GatewayApiGroup gateway.ApiGroup
	TargetApiGroup  target.ApiGroup
	SystemApiGroup  system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
