// 自动生成模板HelpMe
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HelpMe 结构体
// 如果含有time.Time 请自行import time包
type HelpMe struct {
	global.GVA_MODEL
	Title   string `json:"title" form:"title" gorm:"column:title;comment:首页显示标题;"`
	EnTitle string `json:"enTitle" form:"enTitle" gorm:"column:en_title;comment:英文标题;"`
	Icon    string `json:"icon" form:"icon" gorm:"column:icon;comment:图标地址;"`
}
