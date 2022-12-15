// 自动生成模板Complaint
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Complaint 结构体
// 如果含有time.Time 请自行import time包
type Complaint struct {
	global.GVA_MODEL
	Label                 string     `json:"label" form:"label" gorm:"column:label;comment:投诉标签;"`
	OrderNum              string     `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:投诉的相应订单;"`
	Contents              string     `json:"contents" form:"contents" gorm:"column:contents;comment:;"`
	ComplainantId         uint       `json:"complainantId" form:"complainantId" gorm:"column:complainant_id;comment:投诉人的ID;"`
	ByComplainantId       int        `json:"byComplainantId" form:"byComplainantId" gorm:"column:by_complainant_id;comment:被投诉人ID;"`
	Reply                 string     `json:"reply" form:"reply" gorm:"column:reply;comment:回复的信息内容;"`
	Status                *int       `json:"status" form:"status" gorm:"column:status;comment:是否处理完成;"`
	ComplainantUserInfo   MinappUser `gorm:"foreignKey:ComplainantId"`
	ByComplainantUserInfo MinappUser `gorm:"foreignKey:ByComplainantId"`
}
