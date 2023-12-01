package global

import (
	"time"

	"gorm.io/gorm"
)

type BASE_MODEL struct {
	ID        uint32         `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	CreatedBy string         // 创建者
	UpdatedAt time.Time      // 创建时间
	UpdatedBy string         // 更新者
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
