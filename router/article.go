package router

import (
	v1 "self-discipline/api/v1"

	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

func (h *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	var articleApi = v1.ApiGroupApp.ArticleApiGroup.Handler
	ArticleRouter := Router.Group("v1/article")
	{
		ArticleRouter.POST("issue", articleApi.Issue)
		ArticleRouter.POST("list", articleApi.GetList)
		ArticleRouter.POST("give", articleApi.Give)
		ArticleRouter.POST("leave", articleApi.Leave)
	}
	return ArticleRouter
}
