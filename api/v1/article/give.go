package article

import (
	"self-discipline/global"
	articleReq "self-discipline/model/article/request"
	"self-discipline/model/common/request"
	"self-discipline/model/common/response"
	"self-discipline/utils"

	"github.com/gin-gonic/gin"
)

// @Tags Base
// @Summary 点赞
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /v1/article/give [post]
func (h *Handler) Give(ctx *gin.Context) {
	var req articleReq.Give
	_ = ctx.ShouldBind(&req)

	if err := utils.Verify(req, utils.GiveVerify); err != nil {
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
	err := articleService.Give(waitUse.ID, req.ArticleID)
	if err != nil {

	}

	response.Ok(ctx)
}
