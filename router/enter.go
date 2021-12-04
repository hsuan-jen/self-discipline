package router

type RouterGroup struct {
	GatwayRouter
	ArticleRouter
	TargetRouter
	SystemRouter
}

var RouterGroupApp = new(RouterGroup)
