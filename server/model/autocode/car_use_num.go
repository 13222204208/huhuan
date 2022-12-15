package autocode

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type CarUsedNum struct {
	global.GVA_MODEL
	OrderId     uint        `json:"orderId" form:"orderId" gorm:"column:order_id; comment:订单id"`
	Uid         uint        `json:"uid" form:"uid" gorm:"column:uid;comment:使用车辆的ID"`
	Avatar      string      `json:"avatar" form:"avatar" gorm:"column:avatar"`
	HelpCarpool HelpCarpool `gorm:"foreignKey:OrderId"`
}
