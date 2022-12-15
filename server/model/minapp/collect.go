package minapp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

//收藏记录列表
type Collect struct {
	global.GVA_MODEL
	Uid        uint                `json:"uid" gorm:"comment:用户id"`
	Type       uint                `json:"type" gorm:"comment:收藏的类型 二手闲置;default:1"`
	OrderId    uint                `json:"orderId" gorm:"column:order_id;comment:收藏的商品订单ID"`
	Remark     string              `json:"remark" gorm:"comment:备注;default:收藏的商品"`
	Status     uint8               `json:"status" gorm:"comment:收藏的状态;default:1"`
	HelpSecond autocode.HelpSecond `gorm:"foreignKey:OrderId"`
}
