package router

type RouterGroup struct {
	BaseRouter
	ArticleRouter
}

var RouterGroupApp = new(RouterGroup)
