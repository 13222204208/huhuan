// 自动生成模板CreditRating
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CreditRating 结构体
// 如果含有time.Time 请自行import time包
type CreditRating struct {
	global.GVA_MODEL
	Title    string `json:"title" form:"title" gorm:"column:title;comment:信用等级标题;"`
	MinGrade *int   `json:"minGrade" form:"minGrade" gorm:"column:min_grade;comment:最小分数;"`
	MaxGrade *int   `json:"maxGrade" form:"maxGrade" gorm:"column:max_grade;comment:最大分数;"`
	Checked  uint   `json:"checked" form:"checked" gorm:"-"`
}
