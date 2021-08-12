package v1

import (
	"self-discipline/api/v1/article"
	"self-discipline/api/v1/user"
	"self-discipline/api/v1/wechat"
)

type ApiGroup struct {
	UserApiGroup    user.ApiGroup
	WechatApiGroup  wechat.ApiGroup
	ArticleApiGroup article.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
