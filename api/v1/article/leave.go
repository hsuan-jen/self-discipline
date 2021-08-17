package article

import (
	"self-discipline/global"
	"self-discipline/model/article"
	articleReq "self-discipline/model/article/request"
	"self-discipline/model/common/request"
	"self-discipline/model/common/response"
	"self-discipline/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) Leave(ctx *gin.Context) {
	var req articleReq.Leave
	_ = ctx.ShouldBind(&req)

	if err := utils.Verify(req, utils.LeaveVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	claims, ok := ctx.Get("claims")
	if !ok {
		global.LOG.Error("获取用户标识错误，发布失败！")
		response.FailWithMessage("获取用户标识错误，发布失败！", ctx)
		return
	}
	waitUse := claims.(*request.CustomClaims)

	err := articleService.Leave(&article.ArticleLeaves{UserID: waitUse.ID, ArticleID: req.ArticleID, Msg: req.Msg})
	if err != nil {
		global.LOG.Error("留言失败！", zap.Any("err", err))
		response.FailWithMessage("留言失败！", ctx)
		return
	}

	response.Ok(ctx)
}
