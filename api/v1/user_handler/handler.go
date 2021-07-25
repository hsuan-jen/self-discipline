package user_handler

import "github.com/gin-gonic/gin"

var _ Handler = &handler{}

type Handler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type handler struct{}

func New() Handler {
	return &handler{}
}
