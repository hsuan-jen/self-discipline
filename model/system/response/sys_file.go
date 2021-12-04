package response

import "self-discipline/model/system"

type SysFileResponse struct {
	File system.SysFile `json:"file"`
}
