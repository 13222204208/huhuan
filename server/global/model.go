package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID           uint           `gorm:"primarykey"` // 主键ID
	CreatedAt    time.Time      // 创建时间
	UpdatedAt    time.Time      // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
	CreationDate string         `json:"creationDate" form:"creationDate" gorm:"comment:'创建日期';column:creation_date;type:char(30)"`
}
