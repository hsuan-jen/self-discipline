package router

type RouterGroup struct {
	GatwayRouter
	ArticleRouter
	TargetRouter
}

var RouterGroupApp = new(RouterGroup)
