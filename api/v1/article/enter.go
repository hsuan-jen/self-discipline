package article

import (
	"self-discipline/service"
)

type Handler struct{}

type ApiGroup struct {
	Handler
}

var articleService = service.ServiceGroupApp.ArticleServiceGroup
