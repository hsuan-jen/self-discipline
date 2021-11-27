package global

type MODEL struct {
	ID        uint64 `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt int32  `json:"created_at"`           // 创建时间
	UpdatedAt int32  `json:"-"`                    // 更新时间
	DeletedAt int32  `json:"-"`                    // 删除时间
}
