package target

import (
	"self-discipline/model/common/response"
	"self-discipline/utils/validator"

	"github.com/gin-gonic/gin"

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


	
}
