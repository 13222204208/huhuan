// 自动生成模板Category
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Category 结构体
// 如果含有time.Time 请自行import time包
type Category struct {
	global.GVA_MODEL
	Title    string     `json:"title" form:"title" gorm:"column:title;comment:分类名称;"`
	Pid      string     `json:"pid" form:"pid" gorm:"column:pid;comment:父ID;size:8;"`
	Children []Category `json:"children" gorm:"-"`
}
