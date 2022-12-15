package autocode

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Refund struct {
	global.GVA_MODEL
	Uid      uint    `json:"uid" gorm:"column:uid;comment:用户id"`
	Type     uint    `json:"type" gorm:"comment:退款;default:1"`
	OrderNum string  `json:"orderNum" gorm:"column:order_num;comment:退款的订单"`
	Money    float64 `json:"money" gorm:"comment:退款的金额;column:money;"`
	Remark   string  `json:"remark" gorm:"comment:备注"`
	Status   uint8   `json:"status" gorm:"comment:状态 1成功; default:1"`
}
