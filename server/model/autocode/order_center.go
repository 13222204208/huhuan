package autocode

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type OrderCenter struct {
	global.GVA_MODEL
	Uid      uint    `json:"uid" form:"uid" gorm:"column:uid;comment:用户的ID"`
	OrderNum string  `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:订单号"`
	Way      string  `json:"way" form:"way" gorm:"column:way;comment: create 我的发布，take 我的接单，up 我的预约"`
	Status   int     `json:"status" form:"status" gorm:"column:status;default:0"`
	Money    float64 `json:"money" form:"money" gorm:"column:money;type:decimal(10,2);default:0"`
}
