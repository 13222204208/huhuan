// 自动生成模板Area
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Area 结构体
// 如果含有time.Time 请自行import time包
type Area struct {
	global.GVA_MODEL
	Title  string `json:"title" form:"title" gorm:"column:title;comment:地区的名称"`
	Pid    string `json:"pid" form:"pid" gorm:"column:pid;comment:地区的父id;"`
	Status int    `json:"status" form:"status" gorm:"column:status;comment:状态"`
}
