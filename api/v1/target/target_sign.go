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

//获取打卡日志
func (a *TargetSignApi) GetTargetSignList(ctx *gin.Context) {
	var req targetReq.TargetSign
	_ = ctx.ShouldBindJSON(&req)

	if ok, msg := validator.Verify(ctx, &req); !ok {
		response.FailWithMessage(msg, ctx)
		return
	}

	if list, total, err := targetSignService.GetTargetSignList(req); err != nil {
		global.LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", ctx)
	}

}
