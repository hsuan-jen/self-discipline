package article

import (
	"self-discipline/global"
	articleReq "self-discipline/model/article/request"
	"self-discipline/model/common/request"
	"self-discipline/model/common/response"
	"self-discipline/utils/validator"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Accept application/x-www-form-urlencoded
// @Produce  application/json
// @Param data formData articleReq.GetList true "关键词, 页码， 页数大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /v1/article/list [post]
func (h *Handler) GetList(ctx *gin.Context) {
	var req articleReq.GetList
	_ = ctx.ShouldBind(&req)

	if ok, msg := validator.Verify(ctx, &req); !ok {
		response.FailWithMessage(msg, ctx)
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
