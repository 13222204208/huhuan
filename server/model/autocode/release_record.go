// 自动生成模板ReleaseRecord
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ReleaseRecord 结构体
// 如果含有time.Time 请自行import time包
type ReleaseRecord struct {
	global.GVA_MODEL
	Uid       uint       `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;"`
	CouponId  uint       `json:"couponId" form:"couponId" gorm:"column:coupon_id;comment:优惠券id;"`
	Title     string     `json:"title" form:"title" gorm:"column:title;comment:优惠券名称;"`
	Condition float64    `json:"condition" form:"condition" gorm:"column:condition;type:decimal(10,2);comment:优惠券使用条件 0 或空 为不限制;"`
	Start     string     `json:"start" form:"start" gorm:"column:start;comment:优惠券开始使用时间;"`
	Over      string     `json:"over" form:"over" gorm:"column:over;comment:优惠券使用结束时间;"`
	Way       float64    `json:"way" form:"way" gorm:"column:way;type:decimal(10,2);comment:优惠方式立减多少元;"`
	Total     *int       `json:"total" form:"total" gorm:"column:total;comment:优惠券数量;default:1"`
	UsedTotal *int       `json:"used_total" form:"used_total" gorm:"column:used_total;comment:使用的优惠券数量;default:0"`
	Remark    string     `json:"remark" form:"remark" gorm:"column:remark;comment:备注"`
	Status    *int       `json:"status" form:"status" gorm:"column:status;comment:是否使用;default:0"`
	UserInfo  MinappUser `gorm:"foreignKey:Uid"`
}
