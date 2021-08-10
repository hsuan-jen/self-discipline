package article

import "self-discipline/global"

type Articles struct {
	global.MODEL
	UserId  uint64 `json:"user_id"`
	Give    int32  `json:"give"`
	Like    int32  `json:"like"`
	Leave   int32  `json:"leave"`
	Content string `json:"content"`
}
