package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" validate:"required"`                // 页码
	PageSize int `json:"pageSize" form:"pageSize" validate:"required,max=15"` // 每页大小
}

// Find by id structure
type GetById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string // 角色ID
}

type Empty struct{}
