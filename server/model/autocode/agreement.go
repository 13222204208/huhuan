// 自动生成模板Agreement
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Agreement 结构体
// 如果含有time.Time 请自行import time包
type Agreement struct {
	global.GVA_MODEL
	Title    string `json:"title" form:"title" gorm:"column:title;comment:协议标题;"`
	Contents string `json:"contents" form:"contents" gorm:"type:text;column:contents;comment:协议内容;"`
	Status   *int   `json:"status" form:"status" gorm:"column:status;comment:协议状态;size:1;"`
}

// TableName Agreement 表名
func (Agreement) TableName() string {
	return "minapp_agreement"
}
