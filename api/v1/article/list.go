package article

import (
	"self-discipline/global"
	articleReq "self-discipline/model/article/request"
	"self-discipline/model/common/request"
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

	claims, ok := ctx.Get("claims")
	if !ok {
		global.LOG.Error("获取用户标识错误")
		response.FailWithMessage("获取用户标识错误", ctx)
		return
	}
	waitUse := claims.(*request.CustomClaims)

	list, err := articleService.GetList(&req, waitUse.ID)
	if err != nil {
		global.LOG.Error("查询动态列表失败！", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.OkWithData(list, ctx)

}
