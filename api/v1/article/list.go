package article

import (
	"self-discipline/global"
	articleReq "self-discipline/model/article/request"
	"self-discipline/model/common/response"
	"self-discipline/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) GetList(ctx *gin.Context) {
	var req articleReq.GetList
	_ = ctx.ShouldBind(&req)

	if err := utils.Verify(req.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	list, err := articleService.GetList(&req)
	if err != nil {
		global.LOG.Error("查询动态列表失败！", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithData(list, ctx)

}
