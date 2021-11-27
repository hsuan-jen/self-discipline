package target

import (
	"self-discipline/global"
	"self-discipline/model/common/response"
	"self-discipline/utils/validator"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	targetReq "self-discipline/model/target/request"
)

type TargetSignApi struct{}

// @Tags TargetSign
// @Summary 获取打卡日志
// @accept application/json
// @Produce application/json
// @Param data body targetReq.TargetSign true "类型, 页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /target/getTargetSignList [get]
func (a *TargetSignApi) GetTargetSignList(ctx *gin.Context) {
	var req targetReq.TargetSign
	_ = ctx.ShouldBindQuery(&req)

	if ok, msg := validator.Verify(ctx, &req); !ok {
		response.FailWithMessage(msg, ctx)
		return
	}

	list, total, err := targetSignService.GetTargetSignList(req)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取成功", ctx)
}
