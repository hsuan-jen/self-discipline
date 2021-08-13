package article

import (
	"self-discipline/global"
	"self-discipline/model/article"
	articleReq "self-discipline/model/article/request"
	"self-discipline/model/common/response"
	"self-discipline/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData articleReq.Issue true "用户标识, 内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /v1/article/issue [post]
func (h *Handler) Issue(ctx *gin.Context) {
	var req articleReq.Issue
	_ = ctx.ShouldBind(&req)

	if err := utils.Verify(req, utils.IssueVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	article := article.Articles{Content: req.Content, UserId: req.UserId}
	err := articleService.AddArticle(&article)

	if err != nil {
		global.LOG.Error("发布失败!", zap.Any("err", err))
		response.FailWithMessage("发布失败!", ctx)
		return
	}

}
