// 自动生成模板Withdraw
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Withdraw 结构体
// 如果含有time.Time 请自行import time包
type Withdraw struct {
	global.GVA_MODEL
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:提现人用户名称;"`
	UserNum    string     `json:"userNum" form:"userNum" gorm:"column:user_num;comment:;"`
	Phone      string     `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
	Uid        uint       `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;"`
	Money      float64    `json:"money" form:"money" gorm:"column:money;comment:提现的金额;"`
	ApplyTime  string     `json:"applyTime" form:"applyTime" gorm:"column:apply_time;comment:申请提现日期;"`
	Status     uint       `json:"status" form:"status" gorm:"column:status;comment:申请的状态 0 未处理 1 通过，2拒绝;default:0"`
	MinAppUser MinappUser `json:"minAppUser" form:"minAppUser" gorm:"foreignKey:Uid"`
}
