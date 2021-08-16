package article

import (
	articleReq "self-discipline/model/article/request"
	"self-discipline/model/common/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetList(ctx *gin.Context) {
	var req articleReq.GetList
	_ = ctx.ShouldBind(&req)

	list, err := articleService.GetList(&req)
	if err != nil {

	}

	response.OkWithData(list, ctx)

}
