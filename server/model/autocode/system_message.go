// 自动生成模板SystemMessage
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SystemMessage 结构体
// 如果含有time.Time 请自行import time包
type SystemMessage struct {
	global.GVA_MODEL
	Uid      int    `json:"uid" form:"uid" gorm:"column:uid;comment:发送给哪个客户;"`
	Contents string `json:"contents" form:"contents" gorm:"column:contents;comment:消息的内容;"`
	Status   uint   `json:"status" form:"status" gorm:"column:status;comment:消息的状态 0 未读 1 已读;default:0"`
}
