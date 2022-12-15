// 自动生成模板PayCoupon
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PayCoupon 结构体
// 如果含有time.Time 请自行import time包
type PayCoupon struct {
	global.GVA_MODEL
	Type      int     `json:"type" form:"type" gorm:"column:type;comment:赠送的类型 1充值余额赠送，2邀请用户赠送;default:1"`
	CouponId  uint    `json:"couponId" form:"couponId" gorm:"column:coupon_id;comment:优惠券id;"`
	Money     float64 `json:"money" form:"money" gorm:"column:money;type: decimal(10,2);comment:充值金额满多少元送优惠券;default:0"`
	Title     string  `json:"title" form:"title" gorm:"column:title;comment:优惠券名称;"`
	Condition float64 `json:"condition" form:"condition" gorm:"column:condition;type:decimal(10,2);comment:优惠券使用条件 0 或空 为不限制;"`
	Start     string  `json:"start" form:"start" gorm:"column:start;comment:优惠券开始使用时间;"`
	Over      string  `json:"over" form:"over" gorm:"column:over;comment:优惠券使用结束时间;"`
	Way       float64 `json:"way" form:"way" gorm:"column:way;type:decimal(10,2);comment:优惠方式立减多少元;"`
	Total     *int    `json:"total" form:"total" gorm:"column:total;comment:优惠券数量;"`
	Status    *int    `json:"status" form:"status" gorm:"column:status;comment:是否使用;"`
}
