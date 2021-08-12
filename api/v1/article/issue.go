package article

import (
	"self-discipline/model/article"
	articleReq "self-discipline/model/article/request"
	"self-discipline/model/common/response"
	"self-discipline/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Issue(ctx *gin.Context) {
	var req articleReq.Issue
	_ = ctx.ShouldBind(&req)

	if err := utils.Verify(req, utils.IssueVerify); err != nil {
		response.FailWithVerify(err.Error(), ctx)
		return
	}

	article := article.Articles{Content: req.Content, UserId: req.UserId}
	err := articleService.AddArticle(&article)

	if err != nil {
		response.FailWithVerify(err.Error(), ctx)
		return
	}

}
