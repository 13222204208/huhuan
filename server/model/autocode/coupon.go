// 自动生成模板Coupon
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Coupon 结构体
// 如果含有time.Time 请自行import time包
type Coupon struct {
	global.GVA_MODEL
	Title     string  `json:"title" form:"title" gorm:"column:title;comment:优惠券名称;"`
	Condition float64 `json:"condition" form:"condition" gorm:"column:condition;type:decimal(10,2);comment:优惠券使用条件 0 或空 为不限制;"`
	Start     string  `json:"start" form:"start" gorm:"column:start;comment:优惠券开始使用时间;"`
	Way       float64 `json:"way" form:"way" gorm:"column:way;type:decimal(10,2);comment:优惠方式立减多少元;"`
	Over      string  `json:"over" form:"over" gorm:"column:over;comment:优惠券使用结束时间;"`
	Send      int     `json:"send" form:"send" gorm:"column:send;comment:优惠券发放总数;"`
	Sent      int     `json:"sent" form:"sent" gorm:"column:sent;comment:已发放的优惠券数量;"`
	Integral  uint    `json:"integral" form:"integral" gorm:"column:integral;comment:0为不能使用积分兑换;"`
}
