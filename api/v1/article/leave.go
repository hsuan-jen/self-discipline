package article

import (
	"self-discipline/global"
	"self-discipline/model/article"
	articleReq "self-discipline/model/article/request"
	"self-discipline/model/common/request"
	"self-discipline/model/common/response"
	"self-discipline/utils/validator"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 留言
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData articleReq.Leave true "动态标识, 内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /v1/article/leave [post]
func (h *Handler) Leave(ctx *gin.Context) {
	var req articleReq.Leave
	_ = ctx.ShouldBind(&req)

	if ok, msg := validator.Verify(ctx, &req); !ok {
		response.FailWithMessage(msg, ctx)
		return
	}

	claims, ok := ctx.Get("claims")
	if !ok {
		global.LOG.Error("获取用户标识错误！")
		response.FailWithMessage("获取用户标识错误！", ctx)
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

func (h *Handler) LeaveList(ctx *gin.Context) {
	var req articleReq.LeaveList
	_ = ctx.ShouldBind(&req)

	if ok, msg := validator.Verify(ctx, &req); !ok {
		response.FailWithMessage(msg, ctx)
		return
	}

	//err := articleService.LeaveList(req)
}
