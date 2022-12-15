// 自动生成模板MoneyRecord
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MoneyRecord 结构体
// 如果含有time.Time 请自行import time包
type MoneyRecord struct {
	global.GVA_MODEL
	Uid          uint       `json:"uid" form:"uid" gorm:"comment:用户id"`
	Type         uint       `json:"type" form:"type" gorm:"comment:记录类型;comment:1 充值，2 提现 ,3 余额支付的记录,4 收益记录"`
	OrderNum     string     `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:使用余额支付的订单"`
	Money        float64    `json:"money" form:"money" gorm:"comment:金额;column:money;type:decimal(10,2)"`
	Remark       string     `json:"remark" form:"remark" gorm:"comment:备注"`
	Status       uint8      `json:"status" form:"status" gorm:"comment:充值提现记录的状态 1成功; default:0"`
	CreationDate string     `json:"creationDate" form:"creationDate" gorm:"column:creation_date;comment:;"`
	UserInfo     MinappUser `gorm:"foreignKey:Uid"`
}
