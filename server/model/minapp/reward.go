package minapp

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Reward struct {
	global.GVA_MODEL
	OrderType    string  `json:"orderType" form:"orderType" gorm:"column:order_type;comment:订单类型，属于什么订单"`
	OrderId      uint    `json:"orderId" form:"orderId" gorm:"column:order_id;comment:打赏的订单ID"`
	OrderNum     string  `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:订单编号"`
	FromOrderNum string  `json:"fromOrderNum" form:"fromOrderNum" gorm:"column:from_order_num;comment:来自哪个订单"`
	Uid          uint    `json:"uid" form:"uid" gorm:"column:uid;comment:打赏人的ID"`
	Tid          uint    `json:"tid" form:"tid" gorm:"column:tid;default:0;comment:被打赏人的ID"`
	Money        float64 `json:"money" form:"money" gorm:"column:money;type:decimal(10,2);comment:打赏金额"`
	Status       int     `json:"status" form:"status" gorm:"column:status;type:tinyint;comment:订单的状态;default:0"`
}
