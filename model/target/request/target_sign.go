package request

import "self-discipline/model/common/request"

type TargetSign struct {
	Type uint8 `json:"type" form:"type" validate:"required" label:"类型"` // 0推荐 1最新
	request.PageInfo
}
