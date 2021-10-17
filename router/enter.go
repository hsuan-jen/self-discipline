package router

type RouterGroup struct {
	GatwayRouter
	ArticleRouter
}

var RouterGroupApp = new(RouterGroup)
