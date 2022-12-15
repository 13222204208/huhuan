// 自动生成模板CheckMerchant
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CheckMerchant 结构体
// 如果含有time.Time 请自行import time包
type CheckMerchant struct {
	global.GVA_MODEL
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:姓名;"`
	Phone      string     `json:"phone" form:"phone" gorm:"column:phone;comment:联系方式;"`
	Licence    string     `json:"licence" form:"licence" gorm:"column:licence;comment:驾照;"`
	Charter    string     `json:"charter" form:"charter" gorm:"column:charter;comment:营业执照;"`
	Wechat     string     `json:"wechat" form:"wechat" gorm:"column:wechat;comment:微信号"`
	Mail       string     `json:"mail" form:"mail" gorm:"column:mail;comment:邮箱;"`
	Time       string     `json:"time" form:"time" gorm:"column:time;comment:申请日期;"`
	Uid        uint       `json:"uid" form:"uid" gorm:"column:uid;comment:申请人ID;"`
	Status     int        `json:"status" form:"status" gorm:"column:status;comment:申请的状态;size:1;default:1"`
	MinAppUser MinappUser `json:"minAppUser" form:"minAppUser" gorm:"foreignKey:Uid"`
}
